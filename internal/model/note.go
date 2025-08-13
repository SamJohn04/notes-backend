package model

type Note struct {
	Id    int    `json:"id"`
	Owner int    `json:"owner"` // id of owner
	Title string `json:"title"`
	Body  string `json:"body"`
}
