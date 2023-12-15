package main

import (
	"encoding/json"
	"net/http"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	response := GetVersionData()

	// Encode and Return response
	json.NewEncoder(w).Encode(&response)
}
