package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	AgreementId      uint
	ID               uint    `gorm:"primary_key;autoIncrement:true"`
	LenderUserId     uint    `gorm:"not null"`
	BorrowerUserId   uint    `gorm:"not null"`
	LoanAmount       float64 `gorm:"not null"`
	SignedAt         time.Time
	DueAt            *time.Time
	IsApproved       bool          `gorm:"default:false"`
	ContractFile     string        `gorm:"not null"`
	Transactions     []Transaction `gorm:"foreignKey:ContractId"`
	Borrower         User          `gorm:"foreignKey:BorrowerUserId"`
	Lender           User          `gorm:"foreignKey:LenderUserId"`
	Agreement        Agreement     `gorm:"foreignKey:AgreementId"`
	TransactionImage *[]byte       `gorm:"default:null"`
}
