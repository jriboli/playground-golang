package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"my-json-server-api-client/models"
	"net/http"
)

const baseUrl string = "https://my-json-server.typicode.com/jriboli/projects-my-json-server"

func main() {
	log.Println("Launching application...")

	allPosts, err := getAllPosts()
	if err != nil {
		log.Println("Something Broke!")
	}

	log.Println(allPosts)
}

func getAllPosts() ([]models.Post, error) {
	specificUrl := baseUrl + "/posts"

	resp, err := http.Get(specificUrl)
	if err != nil {
		log.Println("Error encountered while calling GET All Posts.")
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error code returned from GET All Posts request - statusCode: %d", resp.StatusCode)
		return nil, nil
	}

	var posts []models.Post
	json.Unmarshal(bodyBytes, &posts)

	return posts, nil
}

func getPostById() {

}

func createPost() {

}

func deletePost() {

}

func getAllComments() {

}

func getAllProfiles() {

}