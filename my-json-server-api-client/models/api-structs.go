package models

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Comment struct {
	Id     int    `json:"id"`
	Body   string `json:"body"`
	PostId int    `json:"post_id"`
}

type Profile struct {
	Name string `json:"name"`
}