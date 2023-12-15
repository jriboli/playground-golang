package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func GetStudyById(w http.ResponseWriter, r *http.Request) {
	// GET Parameters from Request
	params := mux.Vars(r) 
	
	// Make call to ClinicalTrials API
	AddStardard()
	AddQueryTerm(params["studyId"])
	response := baseGetStudy(queryString.String())

	// Encode and Return response
	json.NewEncoder(w).Encode(&response)
}

// func GetStudyByConditions(w http.ResponseWriter, r *http.Request) {
// 	// GET Parameters from Request
// 	params := mux.Vars(r)

// 	// Make call to ClinicalTrials API
// 	queryTerm := strings.Builder{}
// 	for index, condition := range params["conditions"] {
// 		if(index > 0){
// 			queryTerm.WriteString(" OR ")
// 		}
// 		queryTerm.WriteString(condition)
// 	}
// 	AddStardard()
// 	AddQueryTerm(queryTerm.String())
// 	response := baseGetStudy(queryTerm.String())

// 	// Encode and Return response
// 	json.NewEncoder(w).Encode(&response)
// }