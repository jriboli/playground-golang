package main

import (
	"fmt"
	"log"
)

var studiesBaseUrl string = baseUrl + "/studies"

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

func baseGetStudy(queryString string) Studies {
	fmt.Println(queryString)

	url := studiesBaseUrl + queryString

	var studies Studies

	err := GetJson(url, &studies)
	if err != nil {
		log.Fatal(err)
	}

	return studies
}