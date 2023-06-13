package dbmodel

import (
	"gorm.io/gorm"
)

type UserTitle string

const (
	MR  UserTitle = "Mr."
	MS  UserTitle = "Ms."
	MRS UserTitle = "Mrs."
)

//func (self *UserTitle) Scan(value interface{}) error {
//	*self = UserTitle(value.([]byte))
//	return nil
//}
//
//func (self UserTitle) Value() (driver.Value, error) {
//	return string(self), nil
//}

type User struct {
	gorm.Model
	ID         uint           `gorm:"primary_key;autoIncrement:true;not null"`
	Title      UserTitle      `sql:"type:ENUM('Mr.','Ms.','Mrs.')" gorm:"column:title"`
	Username   string         `gorm:"not null"`
	Firstname  string         `gorm:"not null"`
	Lastname   string         `gorm:"not null"`
	Email      string         `gorm:"unique"`
	Password   string         `gorm:"not null"`
	IsLender   bool           `gorm:"not null"`
	Kycs       []Kyc          `gorm:"foreignKey:UserId"`
	Subs       []Subscription `gorm:"foreignKey:UserId"`
	Agreements []Agreement    `gorm:"foreignKey:UserId"`
	Lenders    []Contract     `gorm:"foreignKey:LenderUserId"`
	Borrowers  []Contract     `gorm:"foreignKey:BorrowerUserId"`
	PayChannel string         `gorm:"default:''"`
	PayNumber  string         `gorm:"default:''"`
}
