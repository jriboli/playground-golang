package models

type GetTopStories struct {
	Status      string  `json:"status"`
	Copyright   string  `json:"copyright"`
	Section     string  `json:"section"`
	LastUpdated string  `json:"last_updated"`
	NumResults  int     `json:"num_results"`
	Results     []Story `json:"results"`
}

type Story struct {
	Section    string `json:"section"`
	Subsection string `json:"subsection"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	Url        string `json:"url"`
}