package models

type Role struct {
	Base
	Name    string `gorm:"not null" json:"name"`
	Comment string `json:"comment"`
}
