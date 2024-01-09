package routes

import (
	"github.com/gorilla/mux"

	"restaurant-management/controllers"
)

func AddFoodRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/foods", controllers.GetFoods).Methods("GET")
	foodRouter.HandleFunc("/foods/{food_id}", controllers.GetFood).Methods("GET")
	foodRouter.HandleFunc("/foods", controllers.CreateFood).Methods("POST")
	foodRouter.HandleFunc("/foods/{food_id}", controllers.UpdateFood).Methods("PUT")
	foodRouter.HandleFunc("/foods/{food_id}", controllers.DeleteFood).Methods("DELETE")
}