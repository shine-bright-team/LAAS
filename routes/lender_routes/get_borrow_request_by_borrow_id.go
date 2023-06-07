package lender_routes

import "github.com/gofiber/fiber/v2"

func GetBorrowRequestByBorrowId(c *fiber.Ctx) error {
	return c.SendString("Get Borrow Request by BorrowId")
}
