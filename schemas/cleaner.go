package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	Role string
	Active bool
}

type UserResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Active bool `json:"active"`
}