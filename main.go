package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	go_ora "github.com/sijms/go-ora/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"mti-cm-be-vendor/routes" // Pastikan jalur impor sesuai dengan struktur proyek Anda

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Mendapatkan nilai dari variabel lingkungan
	dbHost := os.Getenv("DB_HOST")
	dbServiceName := os.Getenv("DB_SERVICE_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Construct the Oracle connection string
	connStr := go_ora.BuildUrl(dbHost, 2230, dbServiceName, dbUsername, dbPassword, nil)
	conn, err := sql.Open("oracle", connStr)
	if err != nil {
		log.Fatal("Failed to connect to Oracle", err)
	}
	// check for error
	fmt.Println("Connected to Oracle database!")

	// Mengkonversi koneksi *sql.DB menjadi *gorm.DB
	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize GORM with Oracle connection:", err)
	}

	// Mengatur rute HTTP
	router := routes.SetupRoutes(gormDB)

	// Mengatur Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:7000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// Menjalankan server HTTP
	log.Fatal(router.Run(":7000"))
}
