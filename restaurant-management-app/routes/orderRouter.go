package routes

import (
	"github.com/gorilla/mux"

	"restaurant-management/controllers"
)

func AddOrderRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	foodRouter.HandleFunc("/orders/{order_id}", controllers.GetOrder).Methods("GET")
	foodRouter.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	foodRouter.HandleFunc("/orders/{Order_id}", controllers.UpdateOrder).Methods("PUT")
}