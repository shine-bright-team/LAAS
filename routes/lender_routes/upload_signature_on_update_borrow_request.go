package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
)

type updateLenderPreferenceImageRequest struct {
	ContractId int     `json:"contract_id" validate:"required"`
	Signature  *[]byte `json:"signature_image" validate:"required" `
}

func UploadSignatureOnUpdateBorrowRequest(c *fiber.Ctx) error {
	data := &updateLenderPreferenceImageRequest{}
	userId := c.Locals("userId").(int)
	var contract dbmodel.Contract
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}

	if result := db.DB.Model(&contract).Preload("Agreement").First(&contract, data.ContractId); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	if contract.LenderUserId != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to do this action")
	}

	contract.SignatureImage = data.Signature

	db.DB.Save(&contract)

	return c.Status(fiber.StatusOK).SendString("Image updated successfully")
}
