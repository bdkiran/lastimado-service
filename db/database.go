package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // here
)

//DB is a global variable tha represents our connection to our database.
var DB *sql.DB

func getConnectionString() string {
	log.Println("Creating the postgreSQL connnection string...")
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "hurt")
	return conn
}

//What is the most elegant way of handling the DB?
//Should it be a global varaible that is called?
//Should we create a more complicated object and initialize it??

//SetupConnection establishes a connection to the postgreSQL db.
// After it tests the connection with a ping then returns a pointer to
// the DB object if successful.
func SetupConnection() *sql.DB {
	conn := getConnectionString()

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database.")
	DB = db
	return db
}
