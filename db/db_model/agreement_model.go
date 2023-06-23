package dbmodel

import (
	"gorm.io/gorm"
)

type Agreement struct {
	gorm.Model
	UserId             uint
	ID                 uint       `gorm:"primary_key;autoIncrement:true"`
	LowestAmount       float32    `gorm:"default:0"`
	HighestAmount      float32    `gorm:"default:0"`
	InterestRate       *float32   `gorm:"default:null"`
	DueIn              int32      `gorm:"not null"`
	ActiveAtLeast      *int16     `gorm:"default:null"`
	BaseSalary         *int32     `gorm:"default:null"`
	Addition           string     `gorm:"default:''"`
	Contracts          []Contract `gorm:"foreignKey:AgreementId"`
	IsInterestPerMonth bool       `gorm:"default:false"`
	Owner              User       `gorm:"foreignKey:UserId"`
}
