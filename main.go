package main

import (
	"gofiber-auth/database"
	"gofiber-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	routes.Setup(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
