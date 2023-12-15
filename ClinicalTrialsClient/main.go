package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var client *http.Client
var baseUrl string = "https://clinicaltrials.gov/api/v2"

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// Multiple files
// Instead of 	>> go run main.go
// Use 			>> go run .
// OR 			>> go run *.go
func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	//create router
	r := mux.NewRouter()
	r.HandleFunc("/study", GetStudy).Methods("GET")
	r.HandleFunc("/study/{studyId}", GetStudyById).Methods("GET")

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