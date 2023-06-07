package authroutes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	. "github.com/shine-bright-team/LAAS/v2/db/db_model"
	. "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type registerRequest struct {
	Title     string `json:"title" validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	IsLender  *bool  `json:"isLender" validate:"required"`
}

var userEnum = []string{"Mr.", "Ms.", "Mrs"}

func Register(c *fiber.Ctx) error {
	data := &registerRequest{}
	if err := utils.RequestValidator(c, data); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(*err)
	}
	if !slices.Contains(userEnum, data.Title) {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Title does not matched with \"Mr.\",\"Ms.\",\"Mrs\"")
	}
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error in internal server, please try again later")
	}
	user := &User{
		Title:     UserTitle(data.Title),
		Firstname: data.Firstname,
		Username:  data.Username,
		Email:     data.Email,
		Lastname:  data.Lastname,
		Password:  hashedPassword,
		IsLender:  *data.IsLender,
	}
	result := db.DB.Create(user)
	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			return c.Status(fiber.StatusBadRequest).SendString("This username/firstname/lastname already exist.")
		}
		return c.Status(fiber.ErrBadRequest.Code).SendString("Database Error")
	}
	token, err := utils.SignToken(fmt.Sprint(user.ID), user.IsLender)

	if err != nil {

		return c.Status(fiber.ErrInternalServerError.Code).SendString("Error signing your token, However we've got you registered. please try logging in later.")
	}
	return c.JSON(&LoginResponse{Token: token})
}
