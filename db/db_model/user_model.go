package dbmodel

type User struct {
	Id        uint `gorm:"primary_key;autoIncrement:true"`
	Firstname string
	Lastname  string
	Email     string `gorm:"unique"`
	Password  string
	isLender  bool
}
