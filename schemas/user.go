package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Role      string
	Active    bool
	ImagemUrl string
}
