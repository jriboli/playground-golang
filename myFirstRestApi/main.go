package main

import (
	"log"
	"net/http"

	"github.com/binaryNomad/golang-learning/myFirstRestApi/initializers"
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/myFirstRestApi/controllers"
)

// Runs before Main
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// Init Router
	// := is type inference
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Pet Store Endpoints
	r.HandleFunc("/api/pet_stores", controllers.getPetStores).Methods("GET")
	r.HandleFunc("/api/pet_stores/{id}", getPetStore).Methods("GET")
	r.HandleFunc("/api/pet_stores", createPetStore).Methods("POST")
	r.HandleFunc("/api/pet_stores/{id}", updatePetStores).Methods("PUT")
	r.HandleFunc("/api/pet_stores/{id}", deletePetStores).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}