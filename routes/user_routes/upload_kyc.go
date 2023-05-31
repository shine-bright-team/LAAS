package userroutes

import "github.com/gofiber/fiber/v2"

func UploadKyc(c *fiber.Ctx) error {
	return c.SendString("Upload KYC")
}
