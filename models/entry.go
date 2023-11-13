package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ID             int       `gorm:"primaryKey"`
	AccountId      int       `gorm:"index"`
	TournamentId   int       `gorm:"index"`
	Division       string
	Range          string
	PaymentStatus  string    `gorm:"type:varchar(20)"`
}
