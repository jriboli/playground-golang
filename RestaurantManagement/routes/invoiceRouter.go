package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddInvoiceRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/invoices", controllers.GetInvoice).Methods("GET")
	foodRouter.HandleFunc("/invoices", controllers.AddInvoice).Methods("POST")
}