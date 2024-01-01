package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/database"
	"github.com/binaryNomad/golang-learning/RestaurantManagement/models"
	"github.com/gorilla/mux"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {
	retrievedFoods, err := database.GetFoods()
	if err != nil {
		http.Error(w, "Unable to find Foods", http.StatusNotFound)
		return
	}

	fmt.Printf("Retrieved food: %+v\n", retrievedFoods)

	json.NewEncoder(w).Encode(retrievedFoods)

	// aggregation
	// matchStage	- find records in datasource
	// groupStage 	- group records together
	// projectStage	- how columns/info will be displayed to user
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["food_id"])
	if err != nil {
		fmt.Println("Error: ID is not a number")
		http.Error(w, "Invalid Food Id", http.StatusBadRequest)
		return
	}

	retrievedFood, err := database.GetFood(id)
	if err != nil {
		http.Error(w, "Unable to find Food with ID="+strconv.Itoa(id), http.StatusNotFound)
		return
	}

	fmt.Printf("Retrieved food: %+v\n", retrievedFood)

	json.NewEncoder(w).Encode(retrievedFood)
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
	var f models.Food
	json.NewDecoder(r.Body).Decode(&f)

	f_id, err := database.CreateFood(f)
	if err != nil {
		fmt.Println("Error creating Food. -- " + err.Error())
		http.Error(w, "Error creating Food", http.StatusInternalServerError)
		return
	}

	f.ID = uint(f_id)
	json.NewEncoder(w).Encode(f)
}

func round(num float64) int {

	return 0
}

func toFixed(num float64, precision int) float64 {

	return 0.00
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not yet implemented", http.StatusNotImplemented)
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["food_id"])
	if err != nil {
		fmt.Println("Error: ID is not a number")
		http.Error(w, "Invalid Food Id", http.StatusBadRequest)
		return
	}

	err = database.DeleteFood(id)
	if err != nil {
		http.Error(w, "Unable to find Food with ID="+strconv.Itoa(id), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("Food deleted")
}