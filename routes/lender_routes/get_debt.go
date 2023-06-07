package lender_routes

import "github.com/gofiber/fiber/v2"

func GetDebt(c *fiber.Ctx) error {
	return c.SendString("debt")
}
