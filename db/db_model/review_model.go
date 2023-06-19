package dbmodel

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID             uint `gorm:"primary_key;autoIncrement:true;not null"`
	Score          int16
	ReviewedUserId uint `gorm:"not null"`
	ReviewerUserId uint `gorm:"not null"`
	ReviewFor      User `gorm:"foreignKey:ReviewedUserId"`
	ReviewBy       User `gorm:"foreignKey:ReviewerUserId"`
}
