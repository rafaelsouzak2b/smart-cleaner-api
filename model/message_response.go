package model

import "time"

type MessageResponse struct {
	Id        int64     `json:"id"`
	Message   string    `json:"message"`
	Telefone  string    `json:"telefone"`
	CreatedAt time.Time `json:"created_at"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
}
