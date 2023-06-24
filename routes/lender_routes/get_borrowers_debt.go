package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
)

type contractWithRemaining struct {
	dbmodel.Contract
	Id              uint
	RemainingAmount float64
	Username        string
	Firstname       string
	Lastname        string
}

func GetBorrowersDebt(c *fiber.Ctx) error {

	userId := c.Locals("userId").(int)

	var contracts []contractWithRemaining

	if result := db.DB.Raw("SELECT *, loan_amount - (select COALESCE(sum(transactions.paid_amount),0) from transactions where contract_id = contracts.id) as remaining_amount from contracts join users u on u.id = contracts.borrower_user_id where lender_user_id = ? AND is_approved = true;", userId).Scan(&contracts); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	responses := make([]globalmodels.BorrowRequestResponse, 0)
	for i := range contracts {
		responses = append(responses, globalmodels.BorrowRequestResponse{
			BorrowId:        contracts[i].Id,
			Username:        contracts[i].Username,
			UserId:          contracts[i].BorrowerUserId,
			Firstname:       contracts[i].Firstname,
			Lastname:        contracts[i].Lastname,
			RequestedAmount: contracts[i].LoanAmount,
			RemainingAmount: &contracts[i].RemainingAmount,
			RequestedAt:     contracts[i].CreatedAt,
			DueDate:         contracts[i].DueAt,
			PayChannel:      nil,
			PayNumber:       nil,
			DebtAnalysis:    nil,
			Reviews:         nil,
		})
	}

	return c.JSON(responses)

}
