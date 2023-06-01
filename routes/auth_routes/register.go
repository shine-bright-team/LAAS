package authroutes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	. "github.com/shine-bright-team/LAAS/v2/db/db_model"
	. "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"golang.org/x/exp/slices"
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}
	if !slices.Contains(userEnum, data.Title) {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&ErrorResponse{
			Type:    "Incorrect title",
			Message: "Title is not matched with \"Mr.\",\"Ms.\",\"Mrs\"",
		})
	}
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&ErrorResponse{
			Type:    "Internal server error",
			Message: "There is an error in internal server, please try again later",
		})
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&ErrorResponse{Type: "Database Error", Message: result.Error.Error()})
	}
	token, err := utils.SignToken(fmt.Sprint(user.ID), user.IsLender)

	if err != nil {

		return c.Status(fiber.ErrInternalServerError.Code).JSON(&ErrorResponse{
			Type:    "Internal server error",
			Message: "Error signing your token, However we've got you registered. please try logging in later.",
		})
	}
	return c.JSON(&LoginResponse{Token: token})
}
