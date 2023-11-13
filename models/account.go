package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	ID        int64
	Name      string
	Surname   string
	TelNumber string `gorm:"unique;not null"`
	Gender    string
	Birthdate time.Time
	Password  string `gorm:"not null"`
}