package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")

	
	fmt.Printf("Loaded JWT_SECRET: %q\n", strings.TrimSpace(secret))
}
