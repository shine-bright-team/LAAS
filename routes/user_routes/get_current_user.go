package userroutes

import "github.com/gofiber/fiber/v2"

func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Current user")
}
