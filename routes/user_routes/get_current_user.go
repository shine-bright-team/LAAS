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
		return c.Status(fiber.StatusNotFound).JSON(globalmodels.ErrorResponse{Type: "Unauthorized", Message: "User not found"})
	}
	var kyc dbmodel.Kyc
	dbResult := db.DB.Where("user_id = ?", userId).Last(&kyc)

	if token, err := utils.SignToken(fmt.Sprint(user.ID), user.IsLender); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(globalmodels.ErrorResponse{Type: "Internal server error", Message: "We're having an issue on server, please try again later"})
	} else {
		return c.JSON(globalmodels.UserInfoResponse{
			Username:  user.Username,
			Email:     user.Email,
			Lastname:  user.Lastname,
			Firstname: user.Firstname,
			Id:        fmt.Sprint(user.ID),
			Title:     string(user.Title),
			IsKyc:     !(dbResult.Error != nil || !kyc.IsApproved),
			IsLender:  user.IsLender,
			Token:     token,
		})
	}

}
