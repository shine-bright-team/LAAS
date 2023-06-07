package lender_routes

import "github.com/gofiber/fiber/v2"

func CreateLenderPreference(c *fiber.Ctx) error {
	return c.SendString("Create Lender Preference")
}
