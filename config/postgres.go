package config

import (
	"fmt"

	"github.com/guicazaroto/learning-go/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Environment.PostgresHost, Environment.PostgresPort, Environment.PostgresUser, Environment.PostgresPassword, Environment.PostgresDb)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Errorf("Error connecting to database")
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{})

	if err != nil {
		logger.Errorf("Error migrating database")
		return nil, err
	}

	return db, nil
}
