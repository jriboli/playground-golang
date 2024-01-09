package routes

import (
	"github.com/gorilla/mux"

	"restaurant-management/controllers"
)

func AddUserRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	foodRouter.HandleFunc("/users/{user_id}", controllers.GetUser).Methods("GET")
	foodRouter.HandleFunc("/users/signup", controllers.CreateUser).Methods("POST")
	foodRouter.HandleFunc("/users/login", controllers.AuthenticateUser).Methods("POST")
}