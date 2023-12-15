package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Multiple files
// Instead of 	>> go run main.go
// Use 			>> go run .
// OR 			>> go run *.go
func main() {
	//create router
	r := mux.NewRouter()
	r.HandleFunc("/study", GetStudy).Methods("GET")
	r.HandleFunc("/study/{studyId}", GetStudyById).Methods("GET")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(r)))
}

// Middleware
// To add a header 'Content-Type'
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
