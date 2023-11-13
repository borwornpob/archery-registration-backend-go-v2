package models

import (
	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	ID          int       `gorm:"primaryKey"`
	Name        string
	Description string
	DateRange   string
	Location    string
	Category    string
	Status      string    `gorm:"type:varchar(20)"`
}
