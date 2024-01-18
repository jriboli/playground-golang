package models

type GetMovieReviews struct {
	Status     string   `json:"status"`
	Copyright  string   `json:"copyright"`
	HasMore    bool     `json:"has_more"`
	NumResults int      `json:"num_results"`
	Results    []Review `json:"results"`
}

type Review struct {
	DisplayTitle    string `json:"display_title"`
	MPAARating      string `json:"mpaa_rating"`
	CriticsPick     int    `json:"critics_pick"`
	Byline          string `json:"byline"`
	Headline        string `json:"headline"`
	SummaryShort    string `json:"summary_short"`
	PublicationDate string `json:"publication_date"`
}