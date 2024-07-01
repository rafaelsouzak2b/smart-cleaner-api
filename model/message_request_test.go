package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMessage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		message := MessageRequest{
			Message:  "mensagem teste",
			Telefone: "15165157",
			Nome:     "nome teste",
			Email:    "teste@teste.com",
		}

		err := message.Validate()

		assert.Nil(t, err)
	})
	t.Run("invalid message", func(t *testing.T) {
		message := MessageRequest{
			Message:  "",
			Telefone: "15165157",
			Nome:     "nome teste",
			Email:    "teste@teste.com",
		}

		err := message.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: message (type: string) is required", err.Error())
	})

	t.Run("invalid telefone", func(t *testing.T) {
		message := MessageRequest{
			Message:  "mensagem teste",
			Telefone: "",
			Nome:     "nome teste",
			Email:    "teste@teste.com",
		}

		err := message.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: telefone (type: string) is required", err.Error())
	})

	t.Run("invalid nome", func(t *testing.T) {
		message := MessageRequest{
			Message:  "mensagem teste",
			Telefone: "15165157",
			Nome:     "",
			Email:    "teste@teste.com",
		}

		err := message.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: nome (type: string) is required", err.Error())
	})

	t.Run("invalid email", func(t *testing.T) {
		message := MessageRequest{
			Message:  "mensagem teste",
			Telefone: "15165157",
			Nome:     "nome teste",
			Email:    "",
		}

		err := message.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "param: email (type: string) is required", err.Error())
	})

}
