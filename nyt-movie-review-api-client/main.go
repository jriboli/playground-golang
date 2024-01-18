package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"nyt-movie-review-api-client/models"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Launching application...")

	// Load local config 
	if err := godotenv.Load(); err != nil {
		panic("Failed to load local config!")
	}
	
	reviewResults, err := getMovieReviews()
	if err != nil {
		log.Println("Something Failed.")
	}

	log.Printf("Results:\n %+v", reviewResults)
}

const baseUrl string = "http://api.nytimes.com"


// Grabs the Movie Reviews from New York Times API
// @return: a GetMovieReviews object
func getMovieReviews() (*models.GetMovieReviews, error) {

	// Build URL
	specificUrl := baseUrl + "/svc/movies/v2/reviews/all.json" 

	// Add Auth
	apiKey := os.Getenv("API_KEY")

	queryParams := url.Values{}
	queryParams.Set("api-key", apiKey)

	fullUrl := fmt.Sprintf("%s?%s", specificUrl, queryParams.Encode())

	// Make HTTP call
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Error encountered while calling Get Movie Reviews - ", err)
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// check response
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error returned from Get Movie Reviews- statuscode %d, message %+v", resp.StatusCode, string(bodyBytes))
		return nil, nil
	}

	// Parse response
	var results models.GetMovieReviews
	json.Unmarshal(bodyBytes, &results)

	return &results, nil
}