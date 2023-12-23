package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddTableRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/tables", controllers.GetTables).Methods("GET")
	foodRouter.HandleFunc("/tables/{table_id}", controllers.GetTable).Methods("GET")
	foodRouter.HandleFunc("/tables", controllers.CreateTable).Methods("POST")
	foodRouter.HandleFunc("/tables/{table_id}", controllers.UpdateTable).Methods("PUT")
}