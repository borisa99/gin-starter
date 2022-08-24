package models

type User struct {
	Base
	FirstName    string `gorm:"not null" json:"firstName"`
	LastName     string `gorm:"not null" json:"lastName"`
	Email        string `gorm:"not null;unique" json:"email"`
	PasswordHash string `gorm:"not null" json:"passwordHash"`
	Roles        []UserRole
}
