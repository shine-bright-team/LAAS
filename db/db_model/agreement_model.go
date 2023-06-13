package dbmodel

import (
	"gorm.io/gorm"
)

type Agreement struct {
	gorm.Model
	UserId       uint
	ID           uint       `gorm:"primary_key;autoIncrement:true"`
	InterestRate float32    `gorm:"not null"`
	DueIn        int32      `gorm:"not null"`
	Addition     string     `gorm:"default:''"`
	Contracts    []Contract `gorm:"foreignKey:AgreementId"`
	Owner        User       `gorm:"foreignKey:UserId"`
}
