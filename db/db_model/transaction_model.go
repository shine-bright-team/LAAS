package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ContractId uint
	ID         uint      `gorm:"primary_key;autoIncrement:true"`
	PaidAmount int64     `gorm:"not null"`
	PaidAt     time.Time `gorm:"not null"`
}
