package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the 'PORT' value from the .env file
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not found in the .env file")
	}


	fmt.Println("PORT: ", port)
}