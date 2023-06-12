package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"log"
)

func GetBorrowerRequest(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)

	var contracts []dbmodel.Contract

	if result := db.DB.Model(&dbmodel.Contract{}).Preload("Borrower").Where("lender_user_id = ? AND is_approved = false", uint(userId)); result.Error != nil {
		log.Print(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	borrowRequestResponses := make([]globalmodels.BorrowRequestResponse, 0)

	for i := range contracts {
		borrowRequestResponses = append(borrowRequestResponses, globalmodels.BorrowRequestResponse{
			BorrowId:        contracts[i].ID,
			Username:        contracts[i].Borrower.Username,
			UserId:          contracts[i].Borrower.ID,
			Firstname:       contracts[i].Borrower.Firstname,
			Lastname:        contracts[i].Borrower.Lastname,
			RequestedAmount: contracts[i].LoanAmount,
			RemainingAmount: nil,
			RequestedAt:     contracts[i].CreatedAt,
			DueDate:         nil,
		})
	}

	return c.JSON(borrowRequestResponses)
}
