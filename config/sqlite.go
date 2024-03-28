package config

import (
	"os"

	"github.com/guicazaroto/learning-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/smartCleaner.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating it")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	
	if err != nil {
		logger.Errorf("Error connecting to database")
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Cleaner{})

	if err != nil {
		logger.Errorf("Error migrating database")
		return nil, err
	}

	return db, nil
}