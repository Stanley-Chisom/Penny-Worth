package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Note     string    `json:"note"`
	Date     time.Time `json:"date"`
}
