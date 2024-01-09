package controllers

import (
	"net/http"

	"restaurant-management/models"
)

func GetOrderItems(w http.ResponseWriter, r *http.Request) {

}

func GetOrderItem(w http.ResponseWriter, r *http.Request) {

}

func GetOrderItemsByOrder(w http.ResponseWriter, r *http.Request) {
	// Get order_id

	// Return List of OrderItems, Error

}

// for above
func ItemsByOrder(id string) (OrderItems []models.OrderItem, err error) {
	// Create a slice of OrderItem with one initialized item
	orderItems := []models.OrderItem{{ ID: 1},}

	return orderItems, nil

}

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {

}

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	
}