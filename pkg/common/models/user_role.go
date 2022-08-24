package models

import "github.com/google/uuid"

type UserRole struct {
	Base
	RoleId uuid.UUID `json:"role_id"`
	Role   Role
	UserId uuid.UUID `json:"user_id"`
	User   User
}
