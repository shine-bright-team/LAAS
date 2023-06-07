package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserId    uint
	ID        uint      `gorm:"primary_key;autoIncrement:true"`
	SubsAt    time.Time `gorm:"not null"`
	SubsEndAt time.Time `gorm:"not null"`
}
