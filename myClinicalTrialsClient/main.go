package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client
var baseUrl string = "https://clinicaltrials.gov/api/v2"

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// Multiple files
// Instead of 	>> go run main.go
// Use 			>> go run .
// OR 			>> go run *.go
func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	//GetVersion()
	//GetMetaData()
	GetStudy("NCT06010511")
}