package main

import (
	"api/internal/link"
	"api/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&link.Link{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}
}
