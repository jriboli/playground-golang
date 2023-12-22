package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddOrderRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/endpoint1", controllers.GetOrder()).Methods("GET")
	foodRouter.HandleFunc("/endpoint2", controllers.AddOrder()).Methods("POST")
}