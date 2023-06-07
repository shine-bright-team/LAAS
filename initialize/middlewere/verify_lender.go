package middlewere

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyLender(c *fiber.Ctx) error {
	isLender := c.Locals("isLender").(bool)
	if isLender {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).SendString("You're not a lender")
}
