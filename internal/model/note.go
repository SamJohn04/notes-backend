package model

type Note struct {
	Id    int    `json:"id"`
	Owner string `json:"owner"` // email id of owner
	Title string `json:"title"`
	Body  string `json:"body"`
}
