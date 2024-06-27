package model

import "github.com/guicazaroto/learning-go/util"

type MessageRequest struct {
	Message  string `json:"message"`
	Telefone string `json:"telefone"`
}

func (r *MessageRequest) Validate() error {
	if r.Message == "" {
		return util.ErrParamIsRequired("message", "string")
	}
	if r.Telefone == "" {
		return util.ErrParamIsRequired("telefone", "string")
	}

	return nil
}
