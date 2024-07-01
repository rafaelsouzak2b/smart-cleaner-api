package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToResponse(t *testing.T) {
	t.Run("ToResponse", func(t *testing.T) {
		cleaner := Cleaner{
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
		}

		response := cleaner.ToResponse()

		assert.Equal(t, cleaner.Cidade, response.Cidade)
		assert.Equal(t, cleaner.Descricao, response.Descricao)
		assert.Equal(t, cleaner.UserInfos.Email, response.Email)
		assert.Equal(t, cleaner.Id, response.Id)
		assert.Equal(t, cleaner.UserInfos.ImagemUrl, response.ImagemUrl)
		assert.Equal(t, cleaner.UserInfos.Name, response.Name)
		assert.Equal(t, cleaner.Telefone, response.Telefone)
		assert.Equal(t, cleaner.Uf, response.Uf)
	})

	t.Run("ToResponseMe", func(t *testing.T) {
		cleaner := Cleaner{
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
		}

		response := cleaner.ToResponseMe()

		assert.Equal(t, cleaner.CPF, response.CPF)
		assert.Equal(t, cleaner.Cep, response.Cep)
		assert.Equal(t, cleaner.Cidade, response.Cidade)
		assert.Equal(t, cleaner.DataNascimento, response.DataNascimento)
		assert.Equal(t, cleaner.Descricao, response.Descricao)
		assert.Equal(t, cleaner.UserInfos.Email, response.Email)
		assert.Equal(t, cleaner.Id, response.Id)
		assert.Equal(t, cleaner.UserInfos.ImagemUrl, response.ImagemUrl)
		assert.Equal(t, cleaner.Logradouro, response.Logradouro)
		assert.Equal(t, cleaner.UserInfos.Name, response.Name)
		assert.Equal(t, cleaner.Numero, response.Numero)
		assert.Equal(t, cleaner.Telefone, response.Telefone)
		assert.Equal(t, cleaner.Uf, response.Uf)

	})
}
