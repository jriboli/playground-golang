package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"nyt-books-api-client/models"

	"github.com/joho/godotenv"
)

const baseUrl string = "http://api.nytimes.com/svc/books/v3"

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		// Handle error
		panic("Error loading .env file")
	}

	// Main App
	fmt.Println("Launching application...")

	// Get a List of all Names (Categories)
	getListResp, err := getListNames()
	if err != nil {
		log.Println("Something happened with GetListNames():", err)
		return
	}

	// Wonder if we can use Stream of something relative to filter results
	weeklyUpdatedList := listUpdatedWeekly(getListResp.Results)
	for _, name := range weeklyUpdatedList {
		//fmt.Printf("ListName: %+v",name)
		fmt.Println("Name: " + name.ListName)
		fmt.Println("Updated: " + name.Updated)
	}

	fmt.Println("---------------------------------------------------------------------")

	weeklyUpdatedNotList := listUpdatedNotWeekly(getListResp.Results)
	for _, name := range weeklyUpdatedNotList {
		//fmt.Printf("ListName: %+v",name)
		fmt.Println("Name: " + name.ListName)
		fmt.Println("Updated: " + name.Updated)
	}

	searchListName := getListResp.Results[0].ListNameEncoded
	getBooksResp, err := getBooksByListName(searchListName)
	if err != nil {
		log.Println("Something happened with GetBooksByListName():", err)
	}
	
	fmt.Println("\n\n")
	fmt.Println(getBooksResp)
	//Fatalln() is the equivalent to calling Println() followed by os.Exit(1).
}

// Return an instance of GetListNamesResponse and an error
func getListNames() (*models.GetListNames, error) {
	fmt.Println("\n\n")
	// Build list/names URL
	specificUrl := baseUrl + "/lists/names.json"
	// Grab api-key from the .env file - keeping secret save
	apiKey := os.Getenv("API_KEY")
	//log.Println("The API KEY, grabbed from .env file: \n"+apiKey)

	// Adding Query Parameters
	queryParams := url.Values{}
	queryParams.Set("api-key", apiKey)

	// Construct full URL with parameters
	fullUrl := fmt.Sprintf("%s?%s", specificUrl, queryParams.Encode())
	//log.Println("Fully build URL:\n" + fullUrl)

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Error making GET request:", err)
		return nil, err
	}

	// The defer ensures that the resp.Body.Close() is executed at the end of method. Extemely important
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		log.Println("Response fialed with status code: %d and\nbody: %s\n", resp.StatusCode, bodyBytes)
		return nil, err
	}

	//Convert response body to string
	//bodyString := string(bodyBytes)
	//fmt.Println("API response as string:\n" + bodyString)

	//Convert response body to struct
	var response models.GetListNames
	json.Unmarshal(bodyBytes, &response)
	//fmt.Printf("API response as struct %+v\n", response)

	return &response, nil
}

func getBooksByListName(listName string) (*models.GetBooksByListName, error) {
	fmt.Println("\n\n")
	// Build list/current/{list_name}
	specificUrl := baseUrl + "/lists/current/" + listName

	apiKey := os.Getenv("API_KEY")

	queryParams := url.Values{}
	queryParams.Set("api-key", apiKey)

	fullUrl := fmt.Sprintf("%s?%s", specificUrl, queryParams.Encode())

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Error calling Get Book By List Name:", err)
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Response failed with status code: %d\nbody: %s\n", resp.StatusCode, bodyBytes)
		return nil, err
	}
	//Convert response body to string
	//bodyString := string(bodyBytes)
	//fmt.Println("API response as string:\n" + bodyString)

	var response models.GetBooksByListName
	json.Unmarshal(bodyBytes, &response)
	//log.Printf("API response as struct %+v\n", response)

	return &response, nil
}

func listUpdatedWeekly(names []models.ListName) []models.ListName {
	var results []models.ListName

	for _, name := range names {
		if name.Updated == "WEEKLY" {
			results = append(results, name)
		}
	}

	return results
}

func listUpdatedNotWeekly(names []models.ListName) []models.ListName {
	var results []models.ListName

	for _, name := range names {
		if name.Updated != "WEEKLY" {
			results = append(results, name)
		}
	}

	return results
}