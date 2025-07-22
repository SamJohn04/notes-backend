package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/SamJohn04/notes-backend/internal/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found. Ignore this message if not needed.")
	}
}

func main() {
	app.Run()
}
