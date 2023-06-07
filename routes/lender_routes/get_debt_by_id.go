package lender_routes

import "github.com/gofiber/fiber/v2"

func GetDebtById(c *fiber.Ctx) error {
	return c.SendString("Get Debt by Id")
}
