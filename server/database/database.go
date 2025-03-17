package database

import (
	"fmt"
	"log"
	"pennyWorth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := `host=localhost user=postgres password=Chisom@22 
	dbname=expenses port=5432 sslmode=disable`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database:", db)
	DB = db

	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.User{})
	fmt.Println("Database migration successful")
}
