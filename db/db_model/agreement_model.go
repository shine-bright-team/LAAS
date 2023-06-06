package dbmodel

import (
	"gorm.io/gorm"
)

type Agreement struct {
	gorm.Model
	ID           uint       `gorm:"primary_key;autoIncrement:true"`
	InterestRate int32      `gorm:"not null"`
	DueIn        int32      `gorm:"not null"`
	Addition     string     `gorm:"default:''"`
	Contracts    []Contract `gorm:"foreignKey:AgreementId"`
}
