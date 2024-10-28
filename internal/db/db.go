package db

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Database var
var DB *sql.DB

// Opens Database
func OpenDB(name string) {
	// Open database and store contents in DB
	DB, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}

	// Close Database after use
	defer DB.Close()
}

// Creates table
func CreateTable(name string) {


	//statement, err := DB.Prepare(fmt.Sprintf())
}




