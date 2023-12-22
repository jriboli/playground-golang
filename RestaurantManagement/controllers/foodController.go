package controllers

import (
	"encoding/json"
	"net/http"
)

func GetFood(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{message:No implemented yet}")
}

func AddFood(w http.ResponseWriter, r *http.Request) {
	
}