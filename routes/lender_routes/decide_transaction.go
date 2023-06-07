package lender_routes

import "github.com/gofiber/fiber/v2"

func DecideTransaction(c *fiber.Ctx) error {
	return c.SendString("Decide Transaction by Transaction Id")
}
