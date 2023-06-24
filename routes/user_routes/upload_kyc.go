package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"regexp"
	"time"
)

type kycRequest struct {
	Birthdate string `json:"birthdate" validate:"required"`
	Address   string `json:"address" validate:"required"`
	IdCard    string `json:"id_card" validate:"required,len=13"`
}

func UploadKyc(c *fiber.Ctx) error {
	data := &kycRequest{}
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}
	userId := c.Locals("userId").(int)
	match, _ := regexp.MatchString("(\\d+-\\d+-\\d+)", data.Birthdate)
	birthDate, err := time.Parse("02-01-2006", data.Birthdate)
	if !match || err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Birthdate must be in format dd-mm-yyyy")
	}
	if birthDate.After(time.Now()) {
		return c.Status(fiber.StatusBadRequest).SendString("Birthdate must be in the past")
	}
	kyc := &dbmodel.Kyc{UserId: uint(userId), Address: data.Address, Birthdate: &birthDate, IsApproved: false, IdCard: data.IdCard}
	result := db.DB.Create(kyc)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error in our database")
	}
	return c.JSON(kyc)
}
