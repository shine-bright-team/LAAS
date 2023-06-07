package dbmodel

import (
	"gorm.io/gorm"
)

type PayChannel struct {
	gorm.Model
	UserId  uint
	ID      uint   `gorm:"primary_key;autoIncrement:true"`
	Channel string `gorm:"not null"`
	Number  string `gorm:"not null"`
}
