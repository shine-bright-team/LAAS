package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type getDebtByIdResponse struct {
	DebtDetail   globalmodels.BorrowRequestResponse `json:"debt_detail"`
	Transactions []debtTransaction                  `json:"transactions"`
}

type debtTransaction struct {
	Id           uint      `json:"id"`
	PaidAmount   float64   `json:"paid_amount"`
	PaidAt       time.Time `json:"paid_at"`
	ErrorMessage *string   `json:"error_message"`
	IsApproved   bool      `json:"is_approved"`
	Status       string    `json:"status"`
}

func GetDebtById(c *fiber.Ctx) error {
	debtIdStr := c.Params("debtId")
	userId := c.Locals("userId").(int)

	debtId, err := strconv.Atoi(debtIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid debt id")
	}
	var contract contractWithRemaining

	if result := db.DB.Raw("select *, loan_amount  -  (select COALESCE(sum(paid_amount),0) from transactions where contract_id = contracts.id AND transactions.is_approved = true) as remaining_amount from contracts join users u on u.id = contracts.borrower_user_id where contracts.id = ?;", debtId).Scan(&contract); result.Error != nil {
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

	transactionsResponse := make([]debtTransaction, 0)

	for i := range transactions {
		status := "PENDING"
		if !transactions[i].IsApproved && transactions[i].ErrMessage != nil {
			status = "ERROR"
		}
		if transactions[i].IsApproved {
			status = "SUCCESS"
		}
		transactionsResponse = append(transactionsResponse, debtTransaction{
			Id:           transactions[i].ID,
			PaidAmount:   transactions[i].PaidAmount,
			PaidAt:       transactions[i].PaidAt,
			ErrorMessage: transactions[i].ErrMessage,
			IsApproved:   transactions[i].IsApproved,
			Status:       status,
		})
	}

	response := getDebtByIdResponse{
		DebtDetail: globalmodels.BorrowRequestResponse{
			BorrowId:        contract.Id,
			UserId:          contract.BorrowerUserId,
			Username:        contract.Username,
			Firstname:       contract.Firstname,
			Lastname:        contract.Lastname,
			RequestedAmount: contract.LoanAmount,
			RemainingAmount: &contract.RemainingAmount,
			RequestedAt:     contract.CreatedAt,
			DueDate:         &contract.DueAt,
		},
		Transactions: transactionsResponse,
	}

	return c.JSON(response)
}
