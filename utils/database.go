package utils

import (
	"database/sql"
	"fmt"
)

// Set up Postgres connection
const (
	host     = "localhost"
	port     = 5432
	user     = "your_user_name"
	password = "postgres"
	dbname   = "go_blog_db"
)

// Keep track of the postgres database client
var PSQLDB *sql.DB

// Helper method for connecting to DB
func ConnectDB() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	PSQLDB = db
	return PSQLDB, err
}

// Helper method for retrieving the DB in other files later
func GetDB() *sql.DB {
	return PSQLDB
}
