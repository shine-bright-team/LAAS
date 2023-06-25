package mock

import (
	"errors"
	"github.com/jaswdr/faker"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"log"
	"math"
	"time"
)

func AssignBorrowerRequest(ownerUserId uint) error {
	fake := faker.New()
	var nonLenderUsers []dbmodel.User
	if result := db.DB.Model(&dbmodel.User{}).Where("is_lender", false).Find(&nonLenderUsers); result.Error != nil {
		return result.Error
	}
	var ownerUserAgreement dbmodel.Agreement
	if result := db.DB.Model(&dbmodel.Agreement{}).Where("user_id", ownerUserId).First(&ownerUserAgreement); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return result.Error
	}
	contracts := make([]dbmodel.Contract, 0)
	minLoanAmount := 100.0
	maxLoanAmount := 3000.0
	for i := range nonLenderUsers {
		contracts = append(contracts, dbmodel.Contract{
			AgreementId:    ownerUserAgreement.ID,
			LenderUserId:   ownerUserId,
			BorrowerUserId: nonLenderUsers[i].ID,
			LoanAmount:     math.Round((minLoanAmount+rand.Float64()*(maxLoanAmount-minLoanAmount))*100) / 100,
			SignedAt:       time.Now(),
			DueAt:          nil,
			IsApproved:     false,
			ContractFile:   fake.File().FilenameWithExtension(),
		})
	}
	if result := db.DB.Create(&contracts); result.Error != nil {
		log.Print(result.Error)
		return result.Error
	}
	return nil
}
