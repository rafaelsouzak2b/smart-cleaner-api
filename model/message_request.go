package model

import "github.com/guicazaroto/learning-go/util"

type MessageRequest struct {
	Message  string `json:"message"`
	Telefone string `json:"telefone"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
}

func (r *MessageRequest) Validate() error {
	if r.Message == "" {
		return util.ErrParamIsRequired("message", "string")
	}
	if r.Telefone == "" {
		return util.ErrParamIsRequired("telefone", "string")
	}
	if r.Nome == "" {
		return util.ErrParamIsRequired("nome", "string")
	}
	if r.Email == "" {
		return util.ErrParamIsRequired("email", "string")
	}

	return nil
}
