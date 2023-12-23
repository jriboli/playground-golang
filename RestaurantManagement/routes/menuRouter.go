package routes

import (
	"github.com/gorilla/mux"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/controllers"
)

func AddMenuRoutes(r *mux.Router) {
	foodRouter := r.PathPrefix("/api").Subrouter()

	foodRouter.HandleFunc("/menus", controllers.GetMenus).Methods("GET")
	foodRouter.HandleFunc("/menus/{menu_id}", controllers.GetMenu).Methods("GET")
	foodRouter.HandleFunc("/menus", controllers.CreateMenu).Methods("POST")
	foodRouter.HandleFunc("/menus/{menu_id}", controllers.UpdateMenu).Methods("PUT")
}
