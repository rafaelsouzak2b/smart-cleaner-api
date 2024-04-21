package config

import (
	"fmt"
	"log"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

type environment struct {
	Port             int    `env:"PORT" required:"true"`
	PostgresHost     string `env:"POSTGRES_HOST" required:"true"`
	PostgresPort     int    `env:"POSTGRES_PORT" default:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" required:"true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" required:"true"`
	PostgresDb       string `env:"POSTGRES_DB" required:"true"`
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
	fmt.Println(Environment.PostgresHost)
}
