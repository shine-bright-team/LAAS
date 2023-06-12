package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"gorm.io/gorm"
	"log"
)

type getDebtByIdResponse struct {
	debtDetail   globalmodels.BorrowRequestResponse
	Transactions []dbmodel.Transaction `json:"transactions"`
}

func GetDebtById(c *fiber.Ctx) error {
	debtId := c.Params("debtId")
	userId := c.Locals("userId").(int)

	var contract contractWithRemaining

	if result := db.DB.Raw("select *, loan_amount  -  (select COALESCE(sum(paid_amount),0) from transactions where contract_id = contracts.id AND transactions.is_approved = true) as remaining_amount from contracts join users u on u.id = contracts.borrower_user_id where id = ?;", debtId).Scan(&contract); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Could not find the contract")
		}
		log.Printf("Error: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	if contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to view this contract")
	}

	var transactions []dbmodel.Transaction

	if result := db.DB.Model(&dbmodel.Transaction{}).Where("contract_id = ?", debtId).Find(&transactions); result.Error != nil {
		log.Printf("Error: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	response := getDebtByIdResponse{
		debtDetail: globalmodels.BorrowRequestResponse{
			BorrowId:        contract.ID,
			UserId:          contract.BorrowerUserId,
			Username:        contract.Username,
			Firstname:       contract.Firstname,
			Lastname:        contract.Lastname,
			RequestedAmount: contract.LoanAmount,
			RemainingAmount: &contract.RemainingAmount,
			RequestedAt:     contract.CreatedAt,
			DueDate:         &contract.DueAt,
		},
		Transactions: transactions,
	}

	return c.JSON(response)
}
