package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Kyc struct {
	gorm.Model
	UserId     uint
	ID         uint       `gorm:"primary_key;autoIncrement:true"`
	Name       string     `gorm:"not null"`
	Birthdate  *time.Time `gorm:"not null"`
	Address    string     `gorm:"not null"`
	IdCard     string     `gorm:"not null"`
	IsApproved bool       `gorm:"default:true"`
}
