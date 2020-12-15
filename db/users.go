package db

import (
	"database/sql"
	"log"
)

//User type that is used acrross the whole application
type User struct {
	UserID   int
	Username string
	Password string
	Email    string
}

//GetUsers gets and returns all the users from the database.
func GetUsers() ([]User, error) {
	log.Println("Getting all users from db")
	sqlStatement := `SELECT * FROM users;`

	rows, err := DB.Query(sqlStatement)
	if err != nil {
		log.Println("Error occured when making query.")
		return nil, err
	}

	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.UserID, &u.Username, &u.Password, &u.Email)
		if err != nil {
			log.Println("Error occuredn when scanning rows")
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error occured when checking rows errors")
		return nil, err
	}

	return users, err
}

//GetUser retrieves a user from the database using the userID
func GetUser(userID int) (User, error) {
	log.Printf("Get user query for %d called", userID)
	sqlStatement := `SELECT * FROM users WHERE user_id=$1;`

	row := DB.QueryRow(sqlStatement, userID)
	var u User
	err := row.Scan(&u.UserID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, err
		}
		return u, err
	}
	return u, nil
}

//InsertUser creates a new user in the database
func InsertUser(user User) error {
	log.Println("Insert new user query called")
	sqlStatement :=
		`INSERT INTO users (username, password, email) VALUES ($1, $2, $3) returning user_id;`

	var returnID int
	err := DB.QueryRow(sqlStatement, user.Username, user.Password, user.Email).Scan(&returnID)
	if err != nil {
		return err
	}
	log.Println(returnID)
	return nil
}

//DeleteUser deletes a user entry from the database.
func DeleteUser(userID int) error {
	log.Println("Delete user query called")

	sqlStatement := `DELETE FROM users WHERE user_id=$1;`

	_, err := DB.Exec(sqlStatement, userID)
	if err != nil {
		return err
	}
	return nil

}

//UpdateUser updates a user in the database
//Should this be split across multiple functions?
func UpdateUser(user User) error {
	log.Println("Update user query called")

	sqlStatement := `UPDATE users SET username=$1, password=$2, email=$3 WHERE user_id=$4;`

	_, err := DB.Exec(sqlStatement, user.Username, user.Password, user.Email, user.UserID)
	if err != nil {
		return err
	}
	return nil
}
