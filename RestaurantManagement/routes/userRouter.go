package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddUserRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/users", controllers.GetUser).Methods("GET")
	foodRouter.HandleFunc("/users/{user_id}", controllers.GetUser).Methods("GET")
	foodRouter.HandleFunc("/users/signup", controllers.AddUser).Methods("POST")
	foodRouter.HandleFunc("/users/login", controllers.AuthenticateUser).Methods("POST")
}