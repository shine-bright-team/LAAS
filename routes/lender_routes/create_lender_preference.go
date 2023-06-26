package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/mock"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/gorm"
)

type createLenderRequest struct {
	StartAmount         *float32 `json:"start_amount" validate:"required"`
	EndAmount           float32  `json:"end_amount" validate:"required"`
	InterestRate        *float32 `json:"interest"`
	DueWithIn           int32    `json:"due_with_in" validate:"required"`
	ActiveAtLeast       *int16   `json:"active_at_least"`
	BaseSalary          *int32   `json:"base_salary"`
	AdditionalAgreement *string  `json:"additional_agreement"`
	PaymentChannel      string   `json:"payment_channel" validate:"required"`
	PaymentNumber       string   `json:"payment_number" validate:"required"`
	IsInterestPerMonth  *bool    `json:"is_interest_per_month" validate:"required"`
}

func CreateLenderPreference(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	data := &createLenderRequest{}
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}

	borrowRequest := dbmodel.Agreement{
		UserId:             uint(userId),
		LowestAmount:       *data.StartAmount,
		HighestAmount:      data.EndAmount,
		InterestRate:       data.InterestRate,
		DueIn:              data.DueWithIn,
		ActiveAtLeast:      data.ActiveAtLeast,
		BaseSalary:         data.BaseSalary,
		Addition:           *data.AdditionalAgreement,
		Contracts:          nil,
		IsInterestPerMonth: *data.IsInterestPerMonth,
		Owner:              dbmodel.User{},
	}

	if result := db.DB.Create(&borrowRequest); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).SendString("The preference for this user already exist already exist.")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}
	var user dbmodel.User

	if result := db.DB.Model(&dbmodel.User{}).First(&user, userId); result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	} else {
		user.PayChannel = data.PaymentChannel
		user.PayNumber = data.PaymentNumber
		if updateResult := db.DB.Save(&user); updateResult.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
		}
	}

	if err := mock.AssignBorrowerRequest(uint(userId)); err != nil {
		return err
	}

	return c.JSON(data)
}
