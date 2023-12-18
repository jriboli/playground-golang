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

	createDb()


	//create router
	r := mux.NewRouter()
	r.HandleFunc("/study", GetStudies).Methods("GET")
	r.HandleFunc("/study/{studyId}", GetStudyById).Methods("GET")
	r.HandleFunc("/study/star", GetStarredStudies).Methods("GET")
	r.HandleFunc("/study/{studyId}/star", AddStarToStudy).Methods("POST")
	r.HandleFunc("/study/{studyId}/star", DeleteStarOnStudy).Methods("DELETE")
	r.HandleFunc("/study/{studyId}/decline", DeclineStudy).Methods("POST")
	

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
