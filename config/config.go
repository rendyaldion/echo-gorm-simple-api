package config

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error from load .env")
	}

	return Config{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
}