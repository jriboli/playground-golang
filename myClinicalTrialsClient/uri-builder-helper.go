package main

import (
	"strings"
)

var finalUrl string = ""
var queryString strings.Builder

func AddDefault() {
	AddNewParamToQueryString("format", "json")
}

func AddTotalCount() {
	AddNewParamToQueryString("countTotal", "true")
}

func AddDefaultFields() {
	AddNewParamToQueryString("fields", "NCTId,BriefTitle,OverallStatus,Condition,Phase,StudyType,EnrollmentCount,StudyFirstPostDate,StartDate,LastUpdatePostDate,CentralContactName,CentralContactEMail")
}

func FirstQueryParam() bool {
	return queryString.String() == ""
}

func AddNewParamToQueryString(paramName string, paramValue string) {
	if !FirstQueryParam() {
		queryString.WriteString("&")
	} else {
		queryString.WriteString("?")
	}

	queryString.WriteString(paramName + "=" + paramValue)
}