package lender_routes

import "github.com/gofiber/fiber/v2"

func GetDebtById(c *fiber.Ctx) error {
	//id := c.Params("debtId")
	//if id == ""
	return c.SendString("Get Debt by Id")
}
