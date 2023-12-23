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
		panic(err.Error())
	}

	fmt.Printf("Retrieved food: %+v\n", retrievedFoods)

	json.NewEncoder(w).Encode(retrievedFoods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["food_id"])
	if err != nil {
		fmt.Println("Error: ID is not a number")
	}

	retrievedFood, err := database.GetFood(id)
	if err != nil {
		panic(err.Error())
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
	}

	err = database.DeleteFood(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Food deleted")
}