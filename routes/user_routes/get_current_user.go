package userroutes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/utils"
)

func GetCurrentUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	//isLender := c.Locals("isLender").(bool)
	var user dbmodel.User
	if userResult := db.DB.First(&user, userId); userResult.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found.")
	}
	var kyc dbmodel.Kyc
	dbResult := db.DB.Where("user_id = ?", userId).Last(&kyc)

	var agreement *dbmodel.Agreement

	agreement = nil

	dbAgreement := db.DB.Model(&dbmodel.Agreement{}).Where("user_id = ?", userId).Last(&agreement)

	if token, err := utils.SignToken(fmt.Sprint(user.ID), user.IsLender); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("We're having an issue on server, please try again later")
	} else {
		return c.JSON(globalmodels.UserInfoResponse{
			Username:       user.Username,
			Email:          user.Email,
			Lastname:       user.Lastname,
			Firstname:      user.Firstname,
			Id:             user.ID,
			Title:          string(user.Title),
			IsKyc:          !(dbResult.Error != nil || !kyc.IsApproved),
			IsSetAgreement: !(dbAgreement.Error != nil || agreement.ID == 0),
			IsLender:       user.IsLender,
			Token:          token,
		})
	}

}
