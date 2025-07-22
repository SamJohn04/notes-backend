package main

import (
	"github.com/SamJohn04/notes-backend/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Ignore this message if not needed.")
	}
}

func main() {
	app.Run()
}
