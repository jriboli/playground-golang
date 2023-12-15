package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var studiesBaseUrl string = baseUrl + "/studies"
var baseUrl string = "https://clinicaltrials.gov/api/v2"

var client *http.Client

func GetJson(url string, target interface{}) error {
	client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetMetaData() {
	url := studiesBaseUrl + "/metadata"

	var studyMetaData []MetaData

	err := GetJson(url, &studyMetaData)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Response: %s", studyMetaData)
	}
}

func GetStudyData(queryString string) Studies {
	fmt.Println(queryString)

	url := studiesBaseUrl + queryString

	var studies Studies

	err := GetJson(url, &studies)
	if err != nil {
		log.Fatal(err)
	}

	return studies
}

func GetVersionData() Version {
	url := baseUrl + "/version"

	var version Version

	err := GetJson(url, &version)
	if err != nil {
		fmt.Printf("Error -- %s\n", err.Error())
	} else {
		fmt.Printf("Response: %s", version)
	}

	return version
}
