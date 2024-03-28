package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type UserResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
}