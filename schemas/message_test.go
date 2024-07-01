package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToRespons(t *testing.T) {
	message := Message{
		Id:        1,
		CleanerId: 1,
		Message:   "mensagem teste",
		Telefone:  "12475962",
		Nome:      "nome teste",
		Email:     "teste@teste.com",
		Cleaner: Cleaner{
			Id:     1,
			UserId: 1,
			UserInfos: User{
				Name:      "teste",
				Email:     "teste@teste.com",
				Password:  "cbcjdsb",
				Role:      "cleaner",
				Active:    true,
				ImagemUrl: "ima_url",
			},
			Telefone:       "1254785",
			CPF:            "115165516",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1248522",
			Logradouro:     "rua teste",
			Numero:         "10",
			Cidade:         "Teste",
			Uf:             "tt",
			Descricao:      "descricao teste",
		},
	}

	response := message.ToResponse()

	assert.Equal(t, message.Email, response.Email)
	assert.Equal(t, message.Id, response.Id)
	assert.Equal(t, message.Message, response.Message)
	assert.Equal(t, message.Nome, response.Nome)
	assert.Equal(t, message.Telefone, response.Telefone)
}
