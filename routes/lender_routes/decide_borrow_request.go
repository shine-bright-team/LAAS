package lender_routes

import "github.com/gofiber/fiber/v2"

func DecideBorrowRequest(c *fiber.Ctx) error {
	return c.SendString("Decide to borrow request")
}
