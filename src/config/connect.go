package config

import (
	"Gin/src/initializers"
	"Gin/src/migrations"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	ConnectDB()
}

func ConnectDB() {
	initializers.LoadEnvVariables()
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslmode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslmode, dbTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	AutoMigrate()
}

func AutoMigrate() {
	if err := migrations.Migrate(DB); err != nil {
		log.Fatalf("failed to migrate user table: %v", err)
	}
}
