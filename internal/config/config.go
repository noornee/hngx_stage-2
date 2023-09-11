package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	PORT           string
	CONNECTION_URI string
	DB_NAME        string
}

func Setup() *Configuration {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	config := Configuration{
		PORT:           os.Getenv("PORT"),
		CONNECTION_URI: os.Getenv("CONNECTION_URI"),
		DB_NAME:        os.Getenv("DB_NAME"),
	}

	return &config
}
