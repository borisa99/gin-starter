package models

type Book struct {
	Base
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
