package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"not null,unique"`
	Password  string
	Role      string
	Active    bool
	ImagemUrl string
}
