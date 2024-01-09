package routes

import (
	"github.com/gorilla/mux"

	"restaurant-management/controllers"
)

func AddInvoiceRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")
	foodRouter.HandleFunc("/invoices/{invoice_id}", controllers.GetInvoice).Methods("GET")
	foodRouter.HandleFunc("/invoices", controllers.CreateInvoice).Methods("POST")
	foodRouter.HandleFunc("/invoices/{invoice_id}", controllers.UpdateInvoice).Methods("PUT")
}