package lender_routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/gorm"
	"time"
)

type paymentChannel struct {
	Channel string `json:"channel" validate:"required"`
	Number  string `json:"number" validate:"required"`
}

type createLenderRequest struct {
	StartAmount         float64          `json:"start_amount" validate:"required"`
	EndAmount           float64          `json:"end_amount" validate:"required"`
	InterestRate        float64          `json:"interest" validate:"required"`
	DueWithIn           time.Time        `json:"due_with_in" validate:"required"`
	ActiveAtLeast       int              `json:"active_at_least"`
	BaseSalary          int              `json:"base_salary"`
	AdditionalAgreement string           `json:"additional_agreement"`
	PaymentChannel      []paymentChannel `json:"payment_channel" validate:"required"`
}

// Todo: Connect payment channel to lender preference

func CreateLenderPreference(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	data := &createLenderRequest{}
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}

	var paymentChannels []dbmodel.PayChannel

	for _, channel := range data.PaymentChannel {
		paymentChannels = append(paymentChannels, dbmodel.PayChannel{
			UserId:  uint(userId),
			Channel: channel.Channel,
			Number:  channel.Number,
		})
	}

	if result := db.DB.Create(data); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).SendString("The preference for this user already exist already exist.")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Database Error")
	}

	return c.JSON(data)
}
