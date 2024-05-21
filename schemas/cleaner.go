package schemas

import (
	"github.com/guicazaroto/learning-go/model"
	"gorm.io/gorm"
)

type Cleaner struct {
	gorm.Model
	UserId         int64  `gorm:"not null"`
	UserInfos      User   `gorm:"foreignKey:UserId"`
	Telefone       string `gorm:"size:11;not null"`
	CPF            string `gorm:"size:11;not null"`
	DataNascimento string `gorm:"not null"`
	Cep            string `gorm:"size:8;not null"`
	Logradouro     string `gorm:"size:100;not null"`
	Numero         string `gorm:"not null"`
	Cidade         string `gorm:"size:50;not null"`
	Uf             string `gorm:"size:2;not null"`
	Descricao      string `gorm:"size:200;not null"`
}

func (c *Cleaner) ToResponse() model.CleanerResponse {
	return model.CleanerResponse{
		Id:             int64(c.ID),
		Name:           c.UserInfos.Name,
		Email:          c.UserInfos.Email,
		Active:         c.UserInfos.Active,
		ImagemUrl:      c.UserInfos.ImagemUrl,
		Telefone:       c.Telefone,
		CPF:            c.CPF,
		DataNascimento: c.DataNascimento,
		Cep:            c.CPF,
		Logradouro:     c.Logradouro,
		Numero:         c.Numero,
		Cidade:         c.Cidade,
		Uf:             c.Uf,
		Descricao:      c.Descricao,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}
