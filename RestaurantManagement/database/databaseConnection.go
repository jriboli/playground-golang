package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// The underscore _ is used as a blank identifier. When it is used as the name of a variable or import,
// it essentially tells the Go compiler that you're not going to use that variable or import directly in your code.

var DB *sql.DB

func ConnnectToMySql() {
	dataSourceName := "restaurant_user:restaurant_user@tcp(localhost:3306)/restaurant_management"
    var err error
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        panic(err.Error())
    }

    // Set the maximum number of open connections
    DB.SetMaxOpenConns(10)

    // Set the maximum number of idle connections
    DB.SetMaxIdleConns(5)

    // Check if the connection to the database is successful
    err = DB.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Successfully connected to the database!")
}