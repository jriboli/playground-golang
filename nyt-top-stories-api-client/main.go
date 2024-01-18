package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"nyt-top-stories-api-client/models"

	"github.com/joho/godotenv"
)

const baseUrl string = "http://api.nytimes.com/svc/topstories/v2"

func main() {
	// load local parameters
	if err := godotenv.Load(); err != nil {
		panic("Error pulling local parameters")
	}
	log.Println("Launching application...")
	response, err := get()
	if err != nil {
		log.Print("Error returned from GET() - ", err)
	}

	noPolitics := storiesNotPolitics(response.Results)
	for _, story := range noPolitics {
		log.Println("Title: ", story.Title)
		log.Println("SubSection", story.Subsection)
	}

	fmt.Println(response)
}

func get() (*models.GetTopStories, error) {
	specificUrl := baseUrl + "/home.json"

	apiKey := os.Getenv("API_KEY")

	queryParams := url.Values{}
	queryParams.Set("api-key", apiKey)

	fullUrl := fmt.Sprintf("%s?%s", specificUrl, queryParams.Encode())

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Encountered an error call Get Top Stories: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("API response a byes: %s", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error returned from API request - StatusCode %d, Message %s", resp.StatusCode, bodyBytes)
	}

	var results models.GetTopStories
	json.Unmarshal(bodyBytes, &results)
	//fmt.Printf("API response assing to struct: %+v\n", results)

	return &results, nil
}

func storiesNotPolitics(stories []models.Story) []models.Story {
	var results []models.Story
	for _, story := range stories {
		if story.Subsection != "politics" {
			results = append(results, story)
		}
	}

	return results;
}