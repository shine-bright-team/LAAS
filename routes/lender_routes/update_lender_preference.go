package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/gorm"
)

func UpdateLenderPreference(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	data := &createLenderRequest{}
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}

	agreement := &dbmodel.Agreement{}

	if result := db.DB.Where("user_id = ?", userId).First(agreement); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).SendString("The preference for this user does not exist.")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}

	agreement.InterestRate = data.InterestRate
	agreement.DueIn = data.DueWithIn
	agreement.Addition = *data.AdditionalAgreement
	agreement.IsInterestPerMonth = *data.IsInterestPerMonth
	agreement.LowestAmount = *data.StartAmount
	agreement.HighestAmount = data.EndAmount
	agreement.ActiveAtLeast = data.ActiveAtLeast
	agreement.BaseSalary = data.BaseSalary

	if result := db.DB.Save(agreement); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}
	if result := db.DB.Model(&dbmodel.User{}).Where("id = ?", userId).Update("pay_channel", data.PaymentChannel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}

	if result := db.DB.Model(&dbmodel.User{}).Where("id = ?", userId).Update("pay_number", data.PaymentNumber); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}

	return c.JSON(data)
}
