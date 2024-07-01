package repository

import (
	"github.com/guicazaroto/learning-go/schemas"
	"gorm.io/gorm"
)

type MessageRepository struct {
	Db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return MessageRepository{
		Db: db,
	}
}

func (m *MessageRepository) GetMessagesByCleanerId(cleanerID string) []schemas.Message {
	var messages []schemas.Message
	m.Db.Where("cleaner_id = ?", cleanerID).Find(&messages)

	return messages
}

func (r *MessageRepository) CreateMessage(message *schemas.Message) error {
	if err := r.Db.Create(&message).Error; err != nil {
		return err
	}
	return nil
}
