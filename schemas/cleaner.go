package schemas

import (
	"gorm.io/gorm"
)

type Cleaner struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Bio int `json:"experience"`
	Active bool `json:"active"`
}