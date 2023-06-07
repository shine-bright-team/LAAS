package middlewere

import (
	"github.com/gofiber/fiber/v2"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
)

func VerifyLender(c *fiber.Ctx) error {
	isLender := c.Locals("isLender").(bool)
	if isLender {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(globalmodels.ErrorResponse{Type: "UnAuthorized", Message: "You're not a lender"})
}
