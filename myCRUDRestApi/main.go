package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// This gets executed before Main()
func inti() {
	ConnectToDatabase()
}

func main() {
	// connect to database
	db, err := sql.Open("sql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//create router
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers(db)).Methods("GET")
	r.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	r.HandleFunc("/users", createUser(db)).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(r)))
}

// Middleware 
// To add a header 'Content-Type'
func jsonContentTypeMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

