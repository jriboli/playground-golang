package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddUserRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/endpoint1", controllers.GetUser()).Methods("GET")
	foodRouter.HandleFunc("/endpoint2", controllers.AddUser()).Methods("POST")
}