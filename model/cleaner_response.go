package model

import "time"

type CleanerResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Active         bool      `json:"active"`
	ImagemUrl string `json:"imagem_url"`
	//Telefone       string    `json:"telefone"`
	//CPF            string    `json:"cpf"`
	//DataNascimento string    `json:"data_nascimento"`
	//Cep            string    `json:"cep"`
	//Logradouro     string    `json:"logradouro"`
	//Numero         string    `json:"numero"`
	//Cidade         string    `json:"cidade"`
	//Uf             string    `json:"uf"`
	Descricao string `json:"descricao"`
	//CreatedAt      time.Time `json:"created_at"`
	//UpdatedAt      time.Time `json:"updated_at"`
}

type CleanerMeResponse struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ImagemUrl      string    `json:"imagem_url"`
	Telefone       string    `json:"telefone"`
	CPF            string    `json:"cpf"`
	DataNascimento string    `json:"data_nascimento"`
	Cep            string    `json:"cep"`
	Logradouro     string    `json:"logradouro"`
	Numero         string    `json:"numero"`
	Cidade         string    `json:"cidade"`
	Uf             string    `json:"uf"`
	Descricao      string    `json:"descricao"`
	CreatedAt      time.Time `json:"created_at"`
}
