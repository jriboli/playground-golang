package main

// https://www.youtube.com/watch?v=W7HK2yD0_U0&list=PL5dTjWUk_cPbjazI1vRuTRZi6o5QlVAAR&index=2

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/database"
	"github.com/binaryNomad/golang-learning/RestaurantManagement/middleware"
	"github.com/binaryNomad/golang-learning/RestaurantManagement/routes"

	"github.com/gorilla/mux"
)

var foodDb *sql.DB = database.connnectToMySql();

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Create Router
	r := mux.NewRouter()

	routes.AddFoodRoutes(r)
	routes.AddInvoiceRoutes(r)
	routes.AddMenuRoutes(r)
	routes.AddOrderItemRoutes(r)
	routes.AddOrderRoutes(r)
	routes.AddTableRoutes(r)
	routes.AddUserRoutes(r)

	// Start Server
	log.Fatal(http.ListenAndServe(":"+port, middleware.jsonContentTypeMiddleware(r)))
}

