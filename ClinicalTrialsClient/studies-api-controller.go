package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStudy(w http.ResponseWriter, r *http.Request) {
	// Make call to ClinicalTrials API
	AddStardard()
	AddPageSize(5)
	response := baseGetStudy(queryString.String())

	// Encode and Return response
	json.NewEncoder(w).Encode(&response)
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