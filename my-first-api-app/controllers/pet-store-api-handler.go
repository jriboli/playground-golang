package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"myFirstRestApi/models"
	"github.com/gorilla/mux"
)

// Init books var as a slice Book struct
// slice is a variable length array
var petStores []models.PetStore

// Get All PetStores
// Every function that is a route handler needs to have Request/Response parameters
func getPetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petStores)
}

func getPetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range petStores {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.PetStore{})
}

func createPetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var petStore models.PetStore
	_ = json.NewDecoder(r.Body).Decode(&petStores)
	petStore.Id = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	petStores = append(petStores, petStore)

	json.NewEncoder(w).Encode(petStore)
}

func updatePetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range petStores {
		if item.Id == params["id"] {
			petStores = append(petStores[:index], petStores[index+1:]...)
			var petStore models.PetStore
			_ = json.NewDecoder(r.Body).Decode(&petStore)
			petStore.Id = params["id"]
			petStores = append(petStores, petStore)
			json.NewEncoder(w).Encode(petStore)
			return
		}
	}

	json.NewEncoder(w).Encode(petStores)
}

func deletePetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range petStores {
		if item.Id == params["id"] {
			petStores = append(petStores[:index], petStores[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(petStores)
}