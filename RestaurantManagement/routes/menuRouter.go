package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddMenuRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/menus", controllers.GetMenu).Methods("GET")
	foodRouter.HandleFunc("/menus", controllers.AddMenu).Methods("POST")
}
