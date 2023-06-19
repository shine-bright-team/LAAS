package dbmodel

import (
	"gorm.io/gorm"
)

type InterestCountingType string

const (
	DAY   InterestCountingType = "Day"
	MONTH InterestCountingType = "Month"
)

type Agreement struct {
	gorm.Model
	UserId        uint
	ID            uint                 `gorm:"primary_key;autoIncrement:true"`
	LowestAmount  float64              `gorm:"default:0"`
	HighestAmount float64              `gorm:"default:0"`
	InterestRate  float32              `gorm:"not null"`
	InterestType  InterestCountingType `sql:"type:ENUM('Day','Month')" gorm:"column:interestType"`
	DueIn         int32                `gorm:"not null"`
	Addition      string               `gorm:"default:''"`
	Contracts     []Contract           `gorm:"foreignKey:AgreementId"`
	Owner         User                 `gorm:"foreignKey:UserId"`
}
