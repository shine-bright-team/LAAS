package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	AgreementId    uint
	ID             uint      `gorm:"primary_key;autoIncrement:true"`
	LenderUserId   uint      `gorm:"not null"`
	BorrowerUserId uint      `gorm:"not null"`
	LoanAmount     int64     `gorm:"not null"`
	SignedAt       time.Time `gorm:"not null"`
	DueAt          time.Time
	IsApproved     bool          `gorm:"default:false"`
	ContractFile   string        `gorm:"not null"`
	Transactions   []Transaction `gorm:"foreignKey:ContractId"`
	Borrower       User          `gorm:"foreignKey:BorrowerUserId"`
}
