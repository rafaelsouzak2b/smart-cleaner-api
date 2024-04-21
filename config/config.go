package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostgres()

	if err != nil {
		logger.Errorf("Error initializing db")
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
