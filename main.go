package main

import (
	"fmt"
	"log"
	"os"

	"mti-cm-be-vendor/routes" // Adjust the import path as necessary

	"github.com/cengsin/oracle"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Retrieve database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbService := os.Getenv("DB_SERVICE")

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbService)

	// Connect to the Oracle database using GORM with cengsin/oracle
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Setup HTTP routes
	router := routes.SetupRoutes(db)

	// Setup Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:7000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// Start HTTP server
	log.Fatal(router.Run(":7000"))
}
