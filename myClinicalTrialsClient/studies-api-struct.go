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
	IdentitifcationModule IdentitifcationModule `json:"identificationModule"`
	StatusModule          StatusModule          `json:"statusModule"`
}

// ---------------------------------------------------------------------
// MODULES
// ---------------------------------------------------------------------
type IdentitifcationModule struct {
	NctId      string `json:"nctId"`
	BriefTitle string `json:"briefTitle"`
}

type StatusModule struct {
	OverallStatus string `json:"overallStatus"`
}