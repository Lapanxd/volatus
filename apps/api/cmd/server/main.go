package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lapanxd/volatus-api/internal/middlewares"
	"github.com/lapanxd/volatus-api/internal/models"
	"github.com/lapanxd/volatus-api/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	log.Println("Connected to database")
	return db
}

func main() {
	db := SetupDatabase()

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()

	r.Use(middlewares.ErrorHandlerMiddleware())

	authGroup := r.Group("/auth")

	routes.HealthRoutes(r)
	routes.AuthRoutes(authGroup, db)

	r.Run(":8080")

}
