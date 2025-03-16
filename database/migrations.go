package database

import (
	"fmt"
	"pennyWorth/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Expense{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
	fmt.Println("Database Migrated")
}
