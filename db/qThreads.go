package db

import (
	"database/sql"
	"log"
	"time"
)

//QThread type is used across the whole application to represent question threads
type QThread struct {
	QThreadID int
	Slug      string
	UserID    int
	Title     string
	CreatedAt time.Time
}

//GetQThreads retrieves all the question threads from the database
func GetQThreads() ([]QThread, error) {
	log.Println("Getting all quetion threads from db")
	sqlStatement := `SELECT * FROM qThreads;`

	rows, err := DB.Query(sqlStatement)
	if err != nil {
		log.Println("Error occured when making query")
		return nil, err
	}

	defer rows.Close()

	qThreads := []QThread{}
	for rows.Next() {
		var qThread QThread
		err = rows.Scan(&qThread.QThreadID, &qThread.Slug, &qThread.UserID, &qThread.Title, &qThread.CreatedAt)
		if err != nil {
			log.Println(err)
			log.Println("Error when scaning row into qThread type")
			return nil, err
		}
		qThreads = append(qThreads, qThread)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error occured when checking row errors")
		return nil, err
	}

	return qThreads, err
}

//GetQThreadBySlug retrieves all the question threads from the database using the slug
func GetQThreadBySlug(slug string) (QThread, error) {
	log.Printf("Getting qthread by slug: %s", slug)
	sqlStatement := `SELECT * FROM qThreads WHERE slug=$1;`

	row := DB.QueryRow(sqlStatement, slug)

	var qThread QThread
	err := row.Scan(&qThread.QThreadID, &qThread.Slug, &qThread.UserID, &qThread.Title, &qThread.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return qThread, err
		}
		return qThread, err
	}
	return qThread, nil
}

//GetQThread retrieves a single Question Thread from the database
func GetQThread(qThreadID int) (QThread, error) {
	log.Printf("Get qThread query called for: %d", qThreadID)
	sqlStatment := `SELECT * FROM qThreads WHERE qThread_id=$1;`

	row := DB.QueryRow(sqlStatment, qThreadID)

	var qT QThread
	err := row.Scan(&qT.QThreadID, &qT.UserID, &qT.Title, &qT.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return qT, err
		}
		return qT, err
	}
	return qT, nil
}

//InsertQThread inserts a new Question Thread into the database
func InsertQThread(qThread QThread) error {
	sqlStatement := `INSERT INTO qThreads (user_id, title) VALUES ($1, $2) returning qThread_id;`
	var returnID int
	err := DB.QueryRow(sqlStatement, qThread.UserID, qThread.Title).Scan(&returnID)
	if err != nil {
		return err
	}
	log.Println(returnID)
	return nil
}

//DeleteQThread deletes a Question Thread from the database
func DeleteQThread(qThreadID int) error {
	sqlStatement := `DELETE FROM qThreads WHERE qThread_id=$1`
	_, err := DB.Exec(sqlStatement, qThreadID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateQThread updates the title field of the Question Thread
func UpdateQThread(qThread QThread) error {
	sqlStatement := `UPDATE qThreads SET title=$1 WHERE qThread_id=$2`

	_, err := DB.Exec(sqlStatement, qThread.Title, qThread.QThreadID)
	if err != nil {
		return err
	}
	return nil
}
