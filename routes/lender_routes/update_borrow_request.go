package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
)

type updateBorrowRequest struct {
	ContractId string `json:"contract_id" validate:"required"`
	IsApproved bool   `json:"is_approved" validate:"required"`
}

func UpdateBorrowRequest(c *fiber.Ctx) error {
	data := &updateBorrowRequest{}
	userId := c.Locals("userId").(int)
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}
	var contract dbmodel.Contract

	db.DB.Model(&contract).First(&contract, data.ContractId)

	if contract.IsApproved {
		return c.Status(fiber.StatusBadRequest).SendString("Contract is already approved")
	}

	if contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to do this action")
	}

	if !data.IsApproved {
		db.DB.Delete(&dbmodel.Contract{}, data.ContractId)
		return c.SendString("Decline request")
	} else {
		db.DB.Model(&contract).Where("id", data.ContractId).Update("is_approved", data.IsApproved)
		return c.SendString("Approved request")
	}
}
