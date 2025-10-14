package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	DBSSLMode   string
	JWTSecret   []byte
	JWTDuration time.Duration
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
	DBSSLMode = os.Getenv("DB_SSLMODE")

	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES")) // => 72 * time.Hour
	JWTDuration = duration
}
