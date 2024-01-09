package routes

import (
	"github.com/gorilla/mux"

	"restaurant-management/controllers"
)

func AddOrderItemRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/orderItems", controllers.GetOrderItems).Methods("GET")
	foodRouter.HandleFunc("/orderItems/{item_id}", controllers.GetOrderItem).Methods("GET")
	foodRouter.HandleFunc("/orderItems-order/{order_id}", controllers.GetOrderItemsByOrder).Methods("GET")
	foodRouter.HandleFunc("/orderItems", controllers.CreateOrderItem).Methods("POST")
	foodRouter.HandleFunc("/orderItems/{item_id}", controllers.UpdateOrderItem).Methods("PUT")
}
