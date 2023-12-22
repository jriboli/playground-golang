package database

import (
	"database/sql"
	"fmt"
)

func ConnnectToMySql() (*sql.DB) {
	// Replace with your MySQL database connection details
    connectionString := "user:password@tcp(localhost:3306)/dbname"

    // Open a connection to the MySQL database
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Connected to MySQL database!")

    // Now you can use 'db' to execute SQL queries
    // ...

	return db
}

func ExecuteQuery(db *sql.DB) {
	// Execute a SELECT query
	rows, err := db.Query("SELECT id, name FROM your_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Process the rows
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
}
