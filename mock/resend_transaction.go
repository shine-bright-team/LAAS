package mock

import (
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"time"
)

func ResendTransaction(contractId uint, price float64) error {
	transaction := dbmodel.Transaction{
		ContractId: contractId,
		PaidAmount: price,
		PaidAt:     time.Now(),
		ErrMessage: nil,
		IsApproved: false,
	}
	if result := db.DB.Create(&transaction); result.Error != nil {
		return result.Error
	}
	return nil
}
