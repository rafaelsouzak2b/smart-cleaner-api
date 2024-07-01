package handler

import (
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/stretchr/testify/mock"
)

type mockCleanerRepository struct {
	mock.Mock
}

func (m *mockCleanerRepository) GetCleaners(city string) []schemas.Cleaner {
	args := m.Called(city)

	if cleaners, ok := args.Get(0).([]schemas.Cleaner); ok {
		return cleaners
	}
	return []schemas.Cleaner{}
}

func (m *mockCleanerRepository) GetCleanerById(cleanerID string) *schemas.Cleaner {
	args := m.Called(cleanerID)

	if cleaner, ok := args.Get(0).(*schemas.Cleaner); ok {
		return cleaner
	}
	return nil
}

func (m *mockCleanerRepository) GetCleanerByEmailAndCpf(email, cpf string) int64 {
	args := m.Called(email, cpf)

	if count, ok := args.Get(0).(int64); ok {
		return count
	}
	return 0
}

func (m *mockCleanerRepository) CreateCleaner(cleaner *schemas.Cleaner) error {
	args := m.Called(cleaner)

	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *mockCleanerRepository) SaveCleaner(cleaner *schemas.Cleaner) error {
	args := m.Called(cleaner)

	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *mockCleanerRepository) DeleteCleaner(cleanerID string) error {
	args := m.Called(cleanerID)

	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *mockCleanerRepository) GetCleanerByEmailAndPassword(email, password string) *schemas.Cleaner {
	args := m.Called(email, password)

	if cleaner, ok := args.Get(0).(*schemas.Cleaner); ok {
		return cleaner
	}
	return nil
}

func (m *mockCleanerRepository) UpdateImgUrlCleaner(cleaner *schemas.Cleaner, location string) error {
	return nil
}

type mockMessageRepository struct {
	mock.Mock
}

func (m *mockMessageRepository) GetMessagesByCleanerId(cleanerID string) []schemas.Message {
	args := m.Called(cleanerID)

	if messages, ok := args.Get(0).([]schemas.Message); ok {
		return messages
	}
	return []schemas.Message{}
}

func (m *mockMessageRepository) CreateMessage(message *schemas.Message) error {
	args := m.Called(message)

	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}
