package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCleaner(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.Nil(t, err)
	})

	t.Run("invalid telefone", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: telefone (type: string) is required", err.Error())
	})

	t.Run("invalid cpf", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: cpf (type: string) is required", err.Error())
	})

	t.Run("invalid cep", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: cep (type: string) is required", err.Error())
	})

	t.Run("invalid logradouro", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: logradouro (type: string) is required", err.Error())
	})

	t.Run("invalid cidade", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: cidade (type: string) is required", err.Error())
	})

	t.Run("invalid uf", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: uf (type: string) is required", err.Error())
	})

	t.Run("invalid descricao", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: descricao (type: string) is required", err.Error())
	})

	t.Run("invalid name", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "",
			Email:          "teste@teste.com",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: name (type: string) is required", err.Error())
	})

	t.Run("invalid email", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "",
			Password:       "sax5",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: email (type: string) is required", err.Error())
	})

	t.Run("invalid password", func(t *testing.T) {
		cleaner := CleanerRequest{
			Telefone:       "1655167",
			CPF:            "1458574596",
			DataNascimento: "2024-05-15T19:21:30+03:00",
			Cep:            "1254785",
			Logradouro:     "rua teste",
			Numero:         "125",
			Cidade:         "teste",
			Uf:             "tt",
			Descricao:      "teste descricao",
			Name:           "teste",
			Email:          "teste@teste.com",
			Password:       "",
			Active:         true,
		}

		err := cleaner.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: password (type: string) is required", err.Error())
	})

}
