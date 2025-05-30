package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gofiber-auth/models"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
