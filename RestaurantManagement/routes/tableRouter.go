package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddTableRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/tables", controllers.GetTable).Methods("GET")
	foodRouter.HandleFunc("/tables", controllers.AddTable).Methods("POST")
}