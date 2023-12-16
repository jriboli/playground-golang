package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Filter struct {
	Name string `json:"name"`
	Conditions []string `json:"conditions"`
}

// GET
// https://stackoverflow.com/questions/46045756/retrieve-optional-query-variables-with-gorilla-mux
func GetStudy(w http.ResponseWriter, r *http.Request) {
	// GET query string parameters
	pageSize := r.URL.Query().Get("pageSize")
	recievedData := r.URL.Query()["filter"][0]
	var decodedFilter Filter
	json.Unmarshal([]byte(recievedData), &decodedFilter)
	fmt.Printf("Decoded Filter: %+v\n", decodedFilter)

	// Make call to ClinicalTrials API
	AddStardard()
	AddFilter()

	queryTerm := strings.Builder{}
	for index, condition := range decodedFilter.Conditions {
		if(index > 0){
			queryTerm.WriteString(" OR ")
		}
		queryTerm.WriteString(condition)
	}
	AddQueryTerm(queryTerm.String())

	i, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Fatal(err)
	}
	AddPageSize(i)

	response := GetStudyData(queryString.String())
	var simplifiedResponse []StudyResponse
	simplifiedResponse = populateStudyResponse(response) 

	// Encode and Return response
	json.NewEncoder(w).Encode(&simplifiedResponse)
}

func GetStudyById(w http.ResponseWriter, r *http.Request) {
	// GET Parameters from Request
	params := mux.Vars(r)

	// Make call to ClinicalTrials API
	AddStardard()
	AddQueryTerm(params["studyId"])
	response := GetStudyData(queryString.String())

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

func populateStudyResponse(studies Studies) []StudyResponse {
	var responseHolder []StudyResponse 

	for _, study := range studies.Studies {
		var studyHolder StudyResponse
		studyHolder.BriefSummary = study.ProtocolSection.DescriptionModule.BriefSummary
		studyHolder.BriefTitle = study.ProtocolSection.IdentitifcationModule.BriefTitle
		// CONDITIONS studyHolder.Conditions
		var conditionsHolder []string
		for _, condition := range study.ProtocolSection.ConditionsModule.Conditions {
			conditionsHolder = append(conditionsHolder, condition)
		}
		studyHolder.Conditions = conditionsHolder
		// CONTACTS studyHolder.Contacts
		var contactsHolder []ContactsResponse
		for _, contact := range study.ProtocolSection.ContactsLocationsModule.CentralContacts {
			var contactHolder ContactsResponse
			contactHolder.Email = contact.Email
			contactHolder.Name = contact.Name
			contactHolder.Phone = contact.Phone
			contactHolder.PhoneExt = contact.PhoneExt

			contactsHolder = append(contactsHolder, contactHolder)
		}
		studyHolder.Contacts = contactsHolder
		studyHolder.EligibilityCriteria = study.ProtocolSection.EligibilityModule.EligibilityCriteria
		studyHolder.EligibilityMaxAge = study.ProtocolSection.EligibilityModule.MaximumAge
		studyHolder.EligibilityMinAge = study.ProtocolSection.EligibilityModule.MinimumAge
		studyHolder.EligibilitySex = study.ProtocolSection.EligibilityModule.Sex
		studyHolder.EnrollmentCount = strconv.Itoa(study.ProtocolSection.DesignModule.EnrollmentInfo.Count)
		studyHolder.FirstPostDate = study.ProtocolSection.StatusModule.StudyFirstPostDateStruct.Date
		studyHolder.LastUpdatePostDate = study.ProtocolSection.StatusModule.LastUpdatePostDateStruct.Date
		studyHolder.NCTId = study.ProtocolSection.IdentitifcationModule.NctId
		studyHolder.OrganizationFullName = study.ProtocolSection.IdentitifcationModule.Organization.FullName
		studyHolder.OverallStatus = study.ProtocolSection.StatusModule.OverallStatus
		// PHASES studyHolder.Phases
		var phasesHolder []string
		for _, phase := range study.ProtocolSection.DesignModule.Phases {
			phasesHolder = append(phasesHolder, phase)
		}
		studyHolder.Phases = phasesHolder
		studyHolder.PrimaryCompletionDate = study.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.Date
		studyHolder.StartDate = study.ProtocolSection.StatusModule.StartDateStruct.Date
		studyHolder.StudyType = study.ProtocolSection.DesignModule.StudyType

		responseHolder = append(responseHolder, studyHolder)
	}

	fmt.Println(responseHolder)
	return responseHolder
}
