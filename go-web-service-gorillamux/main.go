package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "About Rocket Raccoon\n")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Accessing Rocket Raccoon\n")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)
	return r
}

// Initiate web server
func main() {
	//fmt.Println("Hello, world.")
	router := router()
	srv := &http.Server{
		Handler: 	router,
		Addr: 		"127.0.0.1:9100",
		WriteTimeout: 	15 * time.Second,
		ReadTimeout:	15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}