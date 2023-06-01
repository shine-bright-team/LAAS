package authroutes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"regexp"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	data := &loginRequest{}

	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}
	var user dbmodel.User
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if emailRegex.MatchString(data.Username) {

		db.DB.Where("email = ?", data.Username).First(&user)

	} else {

		db.DB.Where("username = ?", data.Username).First(&user)

	}

	print(6)

	if utils.CheckPasswordHash(data.Password, user.Password) {
		token, err := utils.SignToken(fmt.Sprint(user.ID), user.IsLender)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(globalmodels.ErrorResponse{
				Type:    "Internal server error",
				Message: "Error signing your token, please try logging in later.",
			})
		}
		return c.JSON(globalmodels.LoginResponse{Token: token})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(globalmodels.ErrorResponse{
			Type:    "Incorrect username or password",
			Message: "Incorrect username (or email) or password",
		})
	}

}
