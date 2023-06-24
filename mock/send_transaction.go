package mock

import (
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"time"
)

func SendTransaction(contract_id string) error {
	var contract dbmodel.Contract
	if result := db.DB.Preload("Borrower").Where("id", contract_id).First(&contract); result.Error != nil {
		return result.Error
	}
	transcations := make([]dbmodel.Transaction, 0)
	money := contract.LoanAmount
	for money > 0 {
		newTransaction := dbmodel.Transaction{
			ContractId: contract.ID,
			PaidAmount: 50,
			PaidAt:     time.Now(),
			ErrMessage: nil,
			IsApproved: false,
		}
		if money < 50 {
			newTransaction.PaidAmount = money
			money = 0
		} else {
			money -= 50
		}
		transcations = append(transcations, newTransaction)
	}
	if result := db.DB.Create(&transcations); result.Error != nil {
		return result.Error
	}
	return nil
}
