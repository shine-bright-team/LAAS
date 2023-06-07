package routes

import "github.com/gofiber/fiber/v2"

func DefaultPage(c *fiber.Ctx) error {
	return c.SendString("Hello, This is LAAS API")
}
