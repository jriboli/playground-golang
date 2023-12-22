package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddMenuRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/endpoint1", controllers.GetMenu()).Methods("GET")
	foodRouter.HandleFunc("/endpoint2", controllers.AddMenu()).Methods("POST")
}
