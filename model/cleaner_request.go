package model

import (
	"github.com/guicazaroto/learning-go/util"
)

type CleanerRequest struct {
	Telefone       string `json:"telefone"`
	CPF            string `json:"cpf"`
	DataNascimento string `json:"data_nascimento"`
	Cep            string `json:"cep"`
	Logradouro     string `json:"logradouro"`
	Numero         string `json:"numero"`
	Cidade         string `json:"cidade"`
	Uf             string `json:"uf"`
	Descricao      string `json:"descricao"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Active         bool   `json:"active"`
}

func (r *CleanerRequest) Validate() error {
	if r.Name == "" {
		return util.ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return util.ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return util.ErrParamIsRequired("password", "string")
	}
	if r.Telefone == "" {
		return util.ErrParamIsRequired("telefone", "string")
	}
	if r.CPF == "" {
		return util.ErrParamIsRequired("cpf", "string")
	}
	if r.Cep == "" {
		return util.ErrParamIsRequired("cep", "string")
	}
	if r.Logradouro == "" {
		return util.ErrParamIsRequired("logradouro", "string")
	}
	if r.Cidade == "" {
		return util.ErrParamIsRequired("cidade", "string")
	}
	if r.Uf == "" {
		return util.ErrParamIsRequired("uf", "string")
	}
	if r.Descricao == "" {
		return util.ErrParamIsRequired("descricao", "string")
	}

	return nil
}
