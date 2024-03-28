package config

import "gorm.io/gorm"

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSqlite()

	if err != nil {
		logger.Errorf("Error initializing sqlite")
		return err
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}