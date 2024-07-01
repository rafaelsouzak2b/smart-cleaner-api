package schemas

import (
	"github.com/guicazaroto/learning-go/model"
	"gorm.io/gorm"
)

type Cleaner struct {
	gorm.Model
	Id             int64
	UserId         int64  `gorm:"not null"`
	UserInfos      User   `gorm:"foreignKey:UserId"`
	Telefone       string `gorm:"not null"`
	CPF            string `gorm:"not null,unique"`
	DataNascimento string `gorm:"not null"`
	Cep            string `gorm:"not null"`
	Logradouro     string `gorm:"size:100;not null"`
	Numero         string `gorm:"not null"`
	Cidade         string `gorm:"size:50;not null"`
	Uf             string `gorm:"size:2;not null"`
	Descricao      string `gorm:"size:200;not null"`
}

func (c *Cleaner) ToResponse() model.CleanerResponse {
	return model.CleanerResponse{
		Id:        c.Id,
		Name:      c.UserInfos.Name,
		Email:     c.UserInfos.Email,
		ImagemUrl: c.UserInfos.ImagemUrl,
		Telefone:  c.Telefone,
		Cidade:    c.Cidade,
		Uf:        c.Uf,
		Descricao: c.Descricao,
	}
}

func (c *Cleaner) ToResponseMe() model.CleanerMeResponse {
	return model.CleanerMeResponse{
		Id:             c.Id,
		Name:           c.UserInfos.Name,
		Email:          c.UserInfos.Email,
		ImagemUrl:      c.UserInfos.ImagemUrl,
		Telefone:       c.Telefone,
		CPF:            c.CPF,
		DataNascimento: c.DataNascimento,
		Cep:            c.Cep,
		Logradouro:     c.Logradouro,
		Numero:         c.Numero,
		Cidade:         c.Cidade,
		Uf:             c.Uf,
		Descricao:      c.Descricao,
		CreatedAt:      c.CreatedAt,
	}
}
