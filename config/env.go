package config

import (
	"log"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

type environment struct {
	PostgresHost        string `env:"POSTGRES_HOST" required:"true"`
	PostgresPort        int    `env:"POSTGRES_PORT" default:"5432"`
	PostgresUser        string `env:"POSTGRES_USER" required:"true"`
	PostgresPassword    string `env:"POSTGRES_PASSWORD" required:"true"`
	PostgresDb          string `env:"POSTGRES_DB" required:"true"`
	AwsRegion           string `env:"AWS_REGION" required:"true"`
	AwsImgProfileBucket string `env:"AWS_IMG_PROFILE_BUCKET" required:"true"`
}

var Environment environment

func InitEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	if err := env.Set(&Environment); err != nil {
		log.Fatal(err)
	}
}
