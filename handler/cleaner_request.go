package handler

import (
	"fmt"
	"time"
)

type CleanerRequest struct {
	Telefone       string    `json:"telefone"`
	CPF            string    `json:"cpf"`
	DataNascimento time.Time `json:"data_nascimento"`
	Cep            string    `json:"cep"`
	Logradouro     string    `json:"logradouro"`
	Numero         int       `json:"numero"`
	Cidade         string    `json:"cidade"`
	Uf             string    `json:"uf"`
	Descricao      string    `json:"descricao"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Active         bool      `json:"active"`
}

func (r *CleanerRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if r.Telefone == "" {
		return errParamIsRequired("telefone", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("cpf", "string")
	}
	// if r.DataNascimento == "" {
	// 	return errParamIsRequired("data_nascimento", "string")
	// }
	if r.Cep == "" {
		return errParamIsRequired("cep", "string")
	}
	if r.Logradouro == "" {
		return errParamIsRequired("logradouro", "string")
	}
	if r.Cidade == "" {
		return errParamIsRequired("cidade", "string")
	}
	if r.Uf == "" {
		return errParamIsRequired("uf", "string")
	}
	if r.Descricao == "" {
		return errParamIsRequired("descricao", "string")
	}

	return nil
}
