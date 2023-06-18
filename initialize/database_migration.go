package initialize

import (
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"gorm.io/gorm"
)

func DatabaseMigration(db_ref *gorm.DB) {
	db_ref.AutoMigrate(&dbmodel.User{}, &dbmodel.Kyc{}, &dbmodel.Subscription{}, &dbmodel.PayChannel{}, &dbmodel.Agreement{}, &dbmodel.Contract{}, &dbmodel.Transaction{}, &dbmodel.Review{})
}
