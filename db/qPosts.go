package db

import (
	"database/sql"
	"log"
	"time"
)

//QPost type is used across the whole application to represents question posts/responses
type QPost struct {
	QPostID   int
	QThreadID int
	UserID    int
	Body      string
	CreatedAt time.Time
}

//GetQPosts retrieves all of the question posts from the database
func GetQPosts() ([]QPost, error) {
	sqlStatement := `SELECT * FROM qPosts;`
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		log.Println("Error occured when retrieving qPosts")
		return nil, err
	}
	defer rows.Close()

	qPosts := []QPost{}
	for rows.Next() {
		var qPost QPost
		err = rows.Scan(&qPost.QPostID, &qPost.QThreadID, &qPost.UserID, &qPost.Body, &qPost.CreatedAt)
		if err != nil {
			log.Println("Error occured when scanning rows into qPost type")
			return nil, err
		}
		qPosts = append(qPosts, qPost)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error occured when checking qPost row errors")
		return nil, err
	}

	return qPosts, err
}

//GetQPost retrieves a single qPost from the database
func GetQPost(qPostID int) (QPost, error) {
	sqlStatement := `SELECT * FROM qPosts WHERE qPost_id=$1;`
	row := DB.QueryRow(sqlStatement, qPostID)

	var qPost QPost
	err := row.Scan(&qPost.QPostID, &qPost.QThreadID, &qPost.UserID, &qPost.Body, &qPost.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return qPost, err
		}
		return qPost, err
	}
	return qPost, nil
}

//InsertQPost inserts a new qPost coloum in the databse
func InsertQPost(qPost QPost) error {
	sqlStatement := `INSERT INTO qPosts (qThread_id, user_id, body) VALUES ($1, $2, $3) RETURNING qPost_id;`

	var returnID int
	err := DB.QueryRow(sqlStatement, qPost.QThreadID, qPost.UserID, qPost.Body).Scan(&returnID)
	if err != nil {
		return err
	}
	log.Println(returnID)
	return nil
}

//UpdateQPost updates the body of the qPost
func UpdateQPost(qPost QPost) error {
	sqlStatement := `UPDATE qPosts SET body=$1 WHERE qPost_id=$2;`
	_, err := DB.Exec(sqlStatement, qPost.Body, qPost.QPostID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteQPost deletes a qPost entry from the database
func DeleteQPost(qPostID int) error {
	sqlStatement := `DELETE FROM qPosts WHERE qPost_id=$1;`
	_, err := DB.Exec(sqlStatement, qPostID)
	if err != nil {
		return err
	}
	return nil
}
