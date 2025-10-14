package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/config"
	"github.com/lapanxd/volatus-api/internal/middleware"
	"github.com/lapanxd/volatus-api/internal/model"
	"github.com/lapanxd/volatus-api/internal/route"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	log.Println("Connected to database")
	return db
}

func main() {
	config.LoadConfig()

	db := SetupDatabase()

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()

	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(middleware.JWTMiddleware())

	authGroup := r.Group("/auth")
	userGroup := r.Group("/user")

	route.HealthRoutes(r)
	route.AuthRoutes(authGroup, db)
	route.UserRoutes(userGroup, db)

	r.Run(":8080")

}
