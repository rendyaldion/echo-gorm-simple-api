package database

import (
	"log"
	"fmt"

	"github.com/rendyaldion/echo-api/config"
	"github.com/rendyaldion/echo-api/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Database struct {
	DB *gorm.DB
}

func InitDatabase() *Database {
	config := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to db")
	}
	
	db.AutoMigrate(&models.User{})

	return &Database{DB:db}
}