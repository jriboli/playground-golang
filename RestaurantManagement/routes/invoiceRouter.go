package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddInvoiceRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/endpoint1", controllers.GetInvoice()).Methods("GET")
	foodRouter.HandleFunc("/endpoint2", controllers.AddInvoice()).Methods("POST")
}