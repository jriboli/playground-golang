package models

// Base response elements
type BaseBooksAPIResponse struct {
	Status       string `json:"status"`
	Copyright    string `json:"copyright"`
	NumResults   int    `json:"num_results"`
	LastModified string `json:"last_modified"`
}

// Method responses
type GetListNames struct {
	// Embedding the Base struct
	Base    BaseBooksAPIResponse
	Results []ListName `json:"results"`
}

type GetBooksByListName struct {
	// Embedding the Base struct
	Base    BaseBooksAPIResponse
	Results ListName `json:"results"`
}

// Embedded objects
type ListName struct {
	ListName            string `json:"list_name"`
	DisplayName         string `json:"display_name"`
	ListNameEncoded     string `json:"list_name_encoded"`
	OldestPublishedDate string `json:"oldest_published_date"`
	NewestPublishedDate string `json:"newest_published_date"`
	Updated             string `json:"updated"`
	BestSellerDate      string `json:"bestseller_date"`
	PublishedDate       string `json:"published_date"`
	Books               []Book `json:"books"`
}

type Book struct {
	Rank        int    `json:"rank"`
	WeeksOnList int    `json:"weeks_on_list"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Title       string `json:"title"`
}