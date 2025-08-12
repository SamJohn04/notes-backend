package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	JWTSecret  string
}

var Cfg Config
var DB *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Ignore this message if not needed.")
	}

	initDB()

	Cfg = Config{
		ServerPort: getEnv("PORT", "8000"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func initDB() {
	// Example: "user:pass@tcp(localhost:3306)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	DB = db
	log.Println("Connected to MariaDB")
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
