package schemas

import (
	"github.com/guicazaroto/learning-go/model"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Id        int64
	CleanerId int64   `gorm:"not null"`
	Cleaner   Cleaner `gorm:"foreignKey:CleanerId"`
	Message   string  `gorm:"size:1000;not null"`
	Telefone  string  `gorm:"not null"`
}

func (m *Message) ToResponse() model.MessageResponse {
	return model.MessageResponse{
		Id:        m.Id,
		Message:   m.Message,
		Telefone:  m.Telefone,
		CreatedAt: m.CreatedAt,
	}
}
