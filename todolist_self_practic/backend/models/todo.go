package models

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
