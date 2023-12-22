package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddOrderItemRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/endpoint1", controllers.GetOrderItem()).Methods("GET")
	foodRouter.HandleFunc("/endpoint2", controllers.AddOrderItem()).Methods("POST")
}
