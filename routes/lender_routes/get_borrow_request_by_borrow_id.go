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
	})
}
