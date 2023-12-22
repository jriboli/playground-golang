package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddFoodRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/foods", controllers.GetFood).Methods("GET")
	foodRouter.HandleFunc("/foods", controllers.AddFood).Methods("POST")
}