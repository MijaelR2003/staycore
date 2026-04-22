package database

import (
	"fmt"
	"log"
	"os"

	"github.com/mijaelr2003/staycore/backend/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Crea las tablas automáticamente
	db.AutoMigrate(
		&models.Guest{},
		&models.Property{},
	)

	log.Println("Database connected successfully")
	DB = db
}