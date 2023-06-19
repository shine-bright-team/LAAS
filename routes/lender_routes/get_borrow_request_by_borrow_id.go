package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"gorm.io/gorm"
	"strconv"
)

type borrowReviewAggregate struct {
	Score       float64
	ReviewCount int
}

type aggratedUserDebt struct {
	Unpaid float64
	Paid   float64
}

func GetBorrowRequestByBorrowId(c *fiber.Ctx) error {

	borrowerIdStr := c.Params("borrowId")
	userId := c.Locals("userId").(int)

	borrowerId, err := strconv.Atoi(borrowerIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid borrow id")
	}

	var contract dbmodel.Contract

	if result := db.DB.Model(&contract).Preload("Borrower").First(&contract, borrowerId); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Contract not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	var aggregatedReview borrowReviewAggregate

	if result := db.DB.Raw("select avg(score) as score, count(*) as review_count from reviews where reviewed_user_id = ?;", contract.BorrowerUserId).Scan(&aggregatedReview); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			aggregatedReview = borrowReviewAggregate{
				Score:       0,
				ReviewCount: 0,
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
		}
	}
	var aggregatedDebt aggratedUserDebt
	if result := db.DB.Raw("with sumOfApprovedTransaction(SummationValue) as (select sum(paid_amount) from transactions join contracts c on c.id = transactions.contract_id where transactions.is_approved = true AND borrower_user_id = ?)select (sum(loan_amount) - SummationValue) as unpaid, SummationValue as paid from contracts,sumOfApprovedTransaction where borrower_user_id = ? group by SummationValue;", contract.BorrowerUserId, contract.BorrowerUserId).Scan(&aggregatedDebt); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			aggregatedDebt = aggratedUserDebt{
				Unpaid: 0,
				Paid:   0,
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
		}
	}

	if contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to do this action")
	}

	if contract.IsApproved {
		return c.Status(fiber.StatusBadRequest).SendString("Contract is already approved")
	}

	return c.JSON(globalmodels.BorrowRequestResponse{
		BorrowId:        contract.ID,
		Username:        contract.Borrower.Username,
		UserId:          contract.BorrowerUserId,
		Firstname:       contract.Borrower.Firstname,
		Lastname:        contract.Borrower.Lastname,
		RequestedAmount: contract.LoanAmount,
		RemainingAmount: nil,
		RequestedAt:     contract.CreatedAt,
		DueDate:         nil,
		PayChannel:      &contract.Borrower.PayChannel,
		PayNumber:       &contract.Borrower.PayNumber,
		DebtAnalysis: &globalmodels.DebtAnalysisResponse{
			Paid:   aggregatedDebt.Paid,
			UnPaid: aggregatedDebt.Unpaid,
		},
		Reviews: &globalmodels.ReviewResponse{
			ReviewAverage: aggregatedReview.Score,
			ReviewCount:   aggregatedReview.ReviewCount,
		},
	})
}
