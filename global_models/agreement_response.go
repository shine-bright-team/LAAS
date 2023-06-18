package globalmodels

type AgreementResponse struct {
	UserId       uint
	ID           uint    `gorm:"primary_key;autoIncrement:true"`
	InterestRate float32 `gorm:"not null"`
	DueIn        int32   `gorm:"not null"`
	Addition     string  `gorm:"default:''"`
}
