package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

// type TransactionStatus string

// const (
// 	WAITING TransactionStatus = "Waiting"
// 	FAILED  TransactionStatus = "Failed"
// 	SUCCESS TransactionStatus = "Success"
// )

type Transaction struct {
	gorm.Model
	ContractId uint
	ID         uint      `gorm:"primary_key;autoIncrement:true"`
	PaidAmount float64   `gorm:"not null"`
	PaidAt     time.Time `gorm:"not null"`
	ErrMessage *string   `gorm:"default:null"`
	IsApproved bool      `gorm:"defaukt:false"`
	// IsApproved TransactionStatus `sql:"type:ENUM('Waiting','Failed','Success')" gorm:"column:isApproved"`
	Contract Contract `gorm:"foreignKey:ContractId"`
}
