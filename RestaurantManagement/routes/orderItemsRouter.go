package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddOrderItemRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/orderItems", controllers.GetOrderItem).Methods("GET")
	foodRouter.HandleFunc("/orderItems", controllers.AddOrderItem).Methods("POST")
}
