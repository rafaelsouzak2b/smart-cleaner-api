package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Cleaner struct {
	gorm.Model
	UserId         int64     `gorm:"not null"`
	UserInfos      User      `gorm:"foreignKey:UserId"`
	Telefone       string    `gorm:"size:11;not null"`
	CPF            string    `gorm:"size:11;not null"`
	DataNascimento time.Time `gorm:"not null"`
	Cep            string    `gorm:"size:8;not null"`
	Logradouro     string    `gorm:"size:100;not null"`
	Numero         int       `gorm:"not null"`
	Cidade         string    `gorm:"size:50;not null"`
	Uf             string    `gorm:"size:2;not null"`
	Descricao      string    `gorm:"size:200;not null"`
}
