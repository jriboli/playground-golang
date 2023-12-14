package main

type MetaData struct {
	AltPieceNames []string `json:"altPieceNames"`
	//Children      []string `json:"children"`
	Definition   string `json:"definition"`
	HistoricOnly string `json:"historicOnly"`
	IndexedOnly  string `json:"indexedOnly"`
	IsEnum       string `json:"isEnum"`
	MaxChars     uint   `json:"maxChars"`
	Name         string `json:"name"`
	Nested       bool   `json:"nested"`
	Piece        string `json:"piece"`
	Rules        string `json:"rules"`
	SourceType   string `json:"sourceType"`
	Synonyms     string `json:"synonyms"`
	Title        string `json:"title"`
	Type         string `json:"type"`
}

// Dont need this
// Use this line --> var studyMetaData []MetaData
// type StudyMetaData struct {
// 	MetaData []MetaData
// }

// ---------------------------------------------------------------------
// STUDIES
// ---------------------------------------------------------------------
type Studies struct {
	TotalCount    uint    `json:"totalCount"`
	Studies       []Study `json:"studies"`
	NextPageToken string  `json:"nextPageToken"`
}

type Study struct {
	ProtocolSection ProtocolSection `json:"protocolSection"`
	HasResults      bool            `json:"hasResults"`
}

// ---------------------------------------------------------------------
// SECTIONS
// ---------------------------------------------------------------------
type ProtocolSection struct {
	IdentitifcationModule      IdentitifcationModule      `json:"identificationModule"`
	StatusModule               StatusModule               `json:"statusModule"`
	SponsorCollaboratorsModule SponsorCollaboratorsModule `json:"sponsorCollaboratorsModule"`
	OversightModule            OversightModule            `json:"oversightModule"`
	DescriptionModule          DescriptionModule          `json:"descriptionModule"`
	ConditionsModule           ConditionsModule           `json:"conditionsModule"`
	DesignModule               DesignModule               `json:"designModule"`
	ArmsInterventionsModule    ArmsInterventionsModule    `json:"armsInterventionsModule"`
	OutcomesModule             OutcomesModule             `json:"outcomesModule"`
	EligibilityModule          EligibilityModule          `json:"eligibilityModule"`
	ContactsLocationsModule    ContactsLocationsModule    `json:"contactsLocationsModule"`
	ReferenceModule            ReferenceModule            `json:"referenceModule"`
	IpdSharingStatementModule  IpdSharingStatementModule  `json:"ipdSharingStatementModule"`
}

// ---------------------------------------------------------------------
// MODULES
// ---------------------------------------------------------------------
type IdentitifcationModule struct {
	NctId           string            `json:"nctId"`
	NctIdAliases    []string          `json:"nctIdAliases"`
	OrgStudIdInfo   OrgStudyIdInfo    `json:"orgStudyIdInfo"`
	SecondaryIdInfo []SecondaryIdInfo `json:"secondaryIdInfos"`
	BriefTitle      string            `json:"briefTitle"`
	OfficialTitle   string            `json:"officialTitle"`
	Acronym         string            `json:"Acronym"`
	Organization    NameStruct        `json:"organization"`
}

type StatusModule struct {
	StatusVerifiedDate          string             `json:"statusVerifiedDate"`
	OverallStatus               string             `json:"overallStatus"`
	LastKnownStatus             string             `json:"lastKnownStatus"`
	DelayedPosting              bool               `json:"delayedPosting"`
	WhyStopped                  string             `json:"whyStopped"`
	ExpandedAccessInfo          ExpandedAccessInfo `json:"expandedAccessInfo"`
	StartDateStruct             DateStruct         `json:"startDateStruct"`
	PrimaryCompletionDateStruct DateStruct         `json:"PrimaryCompletionDateStruct"`
	CompletionDateStruct        DateStruct         `json:"completionDateStruct"`
	StudyFirstSubmitDate        string             `json:"studyFirstSubmitDate"`
	StudyFirstSubmitQcDate      string             `json:"studyFirstSubmitQcDate"`
	StudyFirstPostDateStruct    DateStruct         `json:"studyFirstPostDateStruct"`
	ResultsFirstSubmitDate      string             `json:"resultsFirstSubmitDate"`
	ResultsFirstSumbitQcDate    string             `json:"resultsFirstSumbitQcDate"`
	ResultsFirstPostDateStruct  DateStruct         `json:"resultsFirstPostDateStruct"`
	DispFirstSubmitDate         string             `json:"dispFirstSubmitDate"`
	DispFirstSubmitQcDate       string             `json:"dispFirstSubmitQcDate"`
	DispFirstPostDateStruct     DateStruct         `json:"dispFirstPostDateStruct"`
	LastUpdateSubmitDate        string             `json:"lastUpdateSubmitDate"`
	LastUpdatePostDateStruct    DateStruct         `json:"lastUpdatePostDateStruct"`
}

type SponsorCollaboratorsModule struct {
	ResponsibleParty ResponsibleParty `json:"responsibleParty"`
	LeadSponsor      NameStruct       `json:"leadSponsor"`
	Collaborators    []NameStruct     `json:"collaborators"`
}

type OversightModule struct {
	OversightHasDmc      bool `json:"oversightHasDmc"`
	IsFdaRegulatedDrug   bool `json:"isFdaRegulatedDrug"`
	IsFdaRegulatedDevice bool `json:"isFdaRegulatedDevice"`
	IsUnapprovedDevice   bool `json:"isUnapprovedDevice"`
	IsPpsd               bool `json:"isPpsd"`
	IsUsExport           bool `json:"isUsExport"`
	Fdaaa801Violation    bool `json:"fdaaa801Violation"`
}

type DescriptionModule struct {
	BriefSummary        string `json:"briefSummary"`
	DetailedDescription string `json:"detailedDescription"`
}

type ConditionsModule struct {
	Conditions []string `json:"conditions"`
	Keywords   []string `json:"keywords"`
}

type DesignModule struct {
	StudyType              string              `json:"studyType"`
	NPtrsToThisExpAccNctId uint                `json:"nPtrsToThisExpAccNctId"`
	ExpandedAccessTypes    ExpandedAccessTypes `json:"expandedAccessTypes"`
	PatientRegistry        bool                `json:"patientRegistry"`
	TargetDuration         string              `json:"targetDuration"`
	Phases                 string              `json:"phases"`
	DesignInfo             DesignInfo          `json:"designInfo"`
	BioSpec                BioSpec             `json:"bioSpec"`
	EnrollmentInfo         EnrollmentInfo      `json:"enrollmentInfo"`
}

type ArmsInterventionsModule struct {
	// Empty on purpose - Dont need at the moment
}

type OutcomesModule struct {
	// Empty on purpose - Dont need at the moment
}

type EligibilityModule struct {
	EligibilityCriteria string   `json:"eligibilityCriteria"`
	HealthyVolunteers   bool     `json:"healthyVolunteers"`
	Sex                 string   `json:"sex"`
	GenderBased         bool     `json:"genderBased"`
	GenderDescription   string   `json:"genderDescription"`
	MinimumAge          string   `json:"minimumAge"`
	MaximumAge          string   `json:"maximumAge"`
	StdAge              []string `json:"stdAge"`
	StudyPopulation     string   `json:"studyPopulation"`
	SamplingMethod      string   `json:"samplingMethod"`
}

type ContactsLocationsModule struct {
	CentralContacts  []ContactStruct   `json:"centralContacts"`
	OverallOfficials []OverallOfficial `json:"overallOfficials"`
	Locations        []Location        `json:"locations"`
}

type ReferenceModule struct {
	// Empty on purpose - Dont need at the moment
}

type IpdSharingStatementModule struct {
	// Empty on purpose - Dont need at the moment
}

// ---------------------------------------------------------------------
// SHARED
// ---------------------------------------------------------------------

type NameStruct struct {
	FullName string `json:"fullName"`
	_Class   string `json:"_class"`
}

type DateStruct struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

type ContactStruct struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	PhoneExt string `json:"phoneExt"`
	Email    string `json:"email"`
}

type Location struct{}

// ---------------------------------------------------------------------
// BASE LEVEL
// ---------------------------------------------------------------------

type OrgStudyIdInfo struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Link string `json:"link"`
}

type SecondaryIdInfo struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Link   string `json:"link"`
	Domain string `json:"domain"`
}

type ExpandedAccessInfo struct {
	HasExpandedAccess bool   `json:"hasExpandedAccess"`
	NctId             string `json:"nctId"`
	StatusForNctId    string `json:"statusForNctId"`
}

type ResponsibleParty struct {
	Type                    string `json:"type"`
	InvestigatorFullName    string `json:"investigatorFullName"`
	InvestigatorTitle       string `json:"investigatorTitle"`
	InvestigatorAffiliation string `json:"investigatorAffiliation"`
	OldNameTitle            string `json:"oldNameTitle"`
	OldOrganization         string `json:"oldOrganization"`
}

type ExpandedAccessTypes struct {
	Individual   bool `json:"individual"`
	Intermediate bool `json:"intermediate"`
	Treatment    bool `json:"treatment"`
}

type DesignInfo struct {
	Allocation                   string `json:"allocation"`
	InterventionModel            string `json:"interventionModel"`
	InterventionModelDescription string `json:"interventionModelDescription"`
	PrimaryPurpose               string `json:"primaryPurpose"`
	ObservationalModel           string `json:"observationalModel"`
	TimePerspective              string `json:"timePerspective"`
	MaskingInfo                  string `json:"maskingInfo"`
}

type BioSpec struct {
	Retention   string `json:"retention"`
	Description string `json:"description"`
}

type EnrollmentInfo struct {
	Count uint   `json:"count"`
	Type  string `json:"type"`
}

type OverallOfficial struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
	Role        string `json:"Role"`
}