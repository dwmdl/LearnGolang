package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   DBConfig
	Auth AuthConfig
}

type DBConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error while loading .env file")
	}

	return &Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
