package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"log"
)

func GetAgreemnt(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	var user dbmodel.User
	if userResult := db.DB.First(&user, userId); userResult.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found.")
	}
	var agreement dbmodel.Agreement
	if res := db.DB.Model(&agreement).Where("user_Id = ?", userId).First(&agreement); res.Error != nil {
		log.Print(res.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}
	return c.JSON(globalmodels.AgreementResponse{
		UserId:       agreement.UserId,
		ID:           agreement.ID,
		InterestRate: agreement.InterestRate,
		DueIn:        agreement.DueIn,
		Addition:     agreement.Addition,
	})

}
