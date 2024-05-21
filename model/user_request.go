package model

import (
	"fmt"

	"github.com/guicazaroto/learning-go/util"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Active   bool   `json:"active"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Role == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return util.ErrParamIsRequired("role", "string")
	}
	if r.Name == "" {
		return util.ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return util.ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return util.ErrParamIsRequired("password", "string")
	}

	return nil
}
