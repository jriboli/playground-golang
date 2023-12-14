package main

import "fmt"

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

func GetStudy(studyId string) {
	AddStardard()
	AddQueryTerm(studyId)
	fmt.Println(queryString.String())

	url := studiesBaseUrl + queryString.String()

	var studies Studies

	err := GetJson(url, &studies)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Response: %s", studies)
	}
}