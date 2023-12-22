package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddOrderRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/orders", controllers.GetOrder).Methods("GET")
	foodRouter.HandleFunc("/orders", controllers.AddOrder).Methods("POST")
}