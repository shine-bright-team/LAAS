package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ContractId uint
	ID         uint      `gorm:"primary_key;autoIncrement:true"`
	PaidAmount float64   `gorm:"not null"`
	PaidAt     time.Time `gorm:"not null"`
	IsApproved bool      `gorm:"default:false"`
}
