package main

// https://www.youtube.com/watch?v=W7HK2yD0_U0&list=PL5dTjWUk_cPbjazI1vRuTRZi6o5QlVAAR&index=2

import (
	"log"
	"net/http"
	"os"

	"restaurant-management/database"
	"restaurant-management/middleware"
	"restaurant-management/routes"

	"github.com/gorilla/mux"
)

// OMG and this took forever --- DONT FORGET
// Ensure that the functions you are trying to call are exported (start with an uppercase letter) in their respective packages
//var foodDb *sql.DB = database.ConnnectToMySql();

func main() {
	// Connect to the database
	database.ConnnectToMySql()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Create Router
	r := mux.NewRouter()

	// Add Routes
	routes.AddFoodRoutes(r)
	routes.AddInvoiceRoutes(r)
	routes.AddMenuRoutes(r)
	routes.AddOrderItemRoutes(r)
	routes.AddOrderRoutes(r)
	routes.AddTableRoutes(r)
	routes.AddUserRoutes(r)

	// Start Server
	// OMG and this took forever --- DONT FORGET
	// Ensure that the functions you are trying to call are exported (start with an uppercase letter) in their respective packages
	log.Fatal(http.ListenAndServe(":"+port, middleware.JsonContentTypeMiddleware(r)))
}

