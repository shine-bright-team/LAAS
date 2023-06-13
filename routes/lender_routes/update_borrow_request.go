package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"time"
)

type updateBorrowRequest struct {
	ContractId int  `json:"contract_id" validate:"required"`
	IsApproved bool `json:"is_approved" validate:"required"`
}

// lender/borrower/request

func UpdateBorrowRequest(c *fiber.Ctx) error {
	data := &updateBorrowRequest{}
	userId := c.Locals("userId").(int)
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}
	var contract dbmodel.Contract

	if result := db.DB.Model(&contract).Preload("Agreement").First(&contract, data.ContractId); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	if contract.IsApproved {
		return c.Status(fiber.StatusBadRequest).SendString("Contract is already approved")
	}

	if contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to do this action")
	}

	if contract.Agreement.DueIn == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Due in is not set")
	}

	if !data.IsApproved {
		db.DB.Delete(&dbmodel.Contract{}, data.ContractId)
		return c.SendString("Decline request")
	} else {
		contract.IsApproved = data.IsApproved
		contract.DueAt = time.Now().Add(time.Hour * 24 * 30 * time.Duration(contract.Agreement.DueIn))
		db.DB.Save(&contract)
		//db.DB.Model(&contract).Where("id", data.ContractId).Update("is_approved", data.IsApproved)
		return c.SendString("Approved request")
	}
}
