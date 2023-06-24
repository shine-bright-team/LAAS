package mock

import (
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"time"
)

func SendTransactionWhenRequestAccepted(contractId string) error {
	var contract dbmodel.Contract
	if result := db.DB.Preload("Borrower").Where("id", contractId).First(&contract); result.Error != nil {
		return result.Error
	}
	transactions := make([]dbmodel.Transaction, 0)
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
		transactions = append(transactions, newTransaction)
	}
	if result := db.DB.Create(&transactions); result.Error != nil {
		return result.Error
	}
	return nil
}
