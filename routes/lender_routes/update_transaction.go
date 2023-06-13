package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/gorm"
)

// /lender/debt/transaction

type decideTransactionRequest struct {
	TransactionId int     `json:"transaction_id" validate:"required"`
	IsApproved    *bool   `json:"is_approved" validate:"required"`
	ErrorMessage  *string `json:"error_message"`
}

func UpdateTransaction(c *fiber.Ctx) error {
	data := &decideTransactionRequest{}
	userId := c.Locals("userId").(int)

	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}

	if !*(data.IsApproved) && data.ErrorMessage == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error message is required if transaction is not approved")
	}

	var transaction dbmodel.Transaction

	if result := db.DB.Model(&transaction).Preload("Contract").First(&transaction, data.TransactionId); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Transaction not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}
	if transaction.Contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to do this action")
	}
	if transaction.IsApproved || transaction.ErrMessage != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Transaction is already approved and could not be changed")
	}
	transaction.IsApproved = *data.IsApproved
	if !(*data.IsApproved) {
		transaction.ErrMessage = data.ErrorMessage
	}

	db.DB.Save(&transaction)

	return c.SendString("Transaction is updated successfully")
}
