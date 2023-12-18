package main

import (
	"fmt"
	"strconv"
	"strings"
)

var queryString strings.Builder

// SPECIFIC QUERY PARAMETERS
func AddDefaultParams() {
	AddNewParamToQueryString("format", "json")
}

func AddTotalCount() {
	AddNewParamToQueryString("countTotal", "true")
}

func AddPageSize(pageSize int) {
	if pageSize > 25 {
		pageSize = 25
	}
	AddNewParamToQueryString("pageSize", strconv.Itoa(pageSize))
}

func AddResponseFields() {
	AddNewParamToQueryString("fields", "NCTId,BriefTitle,OverallStatus,Condition,Phase,StudyType,EnrollmentCount,"+
		"StudyFirstPostDate,StartDate,LastUpdatePostDate,CentralContactName,CentralContactEMail,OrgFullName,BriefSummary,EligibilityCriteria,Sex,MinimumAge,MaximumAge")
}

func AddQueryTerm(conditions string) {
	AddNewParamToQueryString("query.term", conditions)
}

func AddFilter() {
	AddNewParamToQueryString("filter.advanced", "(((AREA[Phase] \"Phase 2\" OR AREA[Phase] \"Phase 3\" OR AREA[Phase] \"Phase 4\")"+
		"AND (AREA[LeadSponsorClass] Industry OR AREA[CollaboratorClass] Industry) NOT (AREA[CentralContactEMail]MISSING)))")
}

func AddFilterOverallStatus() {
	AddNewParamToQueryString("filter.overallStatus", "NOT_YET_RECRUITING,RECRUITING")
}

func AddSorting(sortBy string) {
	// String format needs to be -- {field name}:{asc/desc}
	AddNewParamToQueryString("sort", sortBy)
}

// BASIC SETUPS
func AddStardard() {
	AddDefaultParams()
	AddTotalCount()
	AddResponseFields()
}

// HELPER FUNCTIONS
func FirstQueryParam() bool {
	return queryString.String() == ""
}

func AddNewParamToQueryString(paramName string, paramValue string) {
	if !FirstQueryParam() {
		queryString.WriteString("&")
	} else {
		queryString.WriteString("?")
	}

	queryString.WriteString(paramName + "=" + BasicHtmlEncode(paramValue))
}

func BasicHtmlEncode(html string) string {
	html = strings.Replace(html, " ", "%20", -1)
	html = strings.ReplaceAll(html, "\"", "%22")

	fmt.Println("Final query: " + html)
	return html
}
