package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func RequestValidator(c *fiber.Ctx, data interface{}) *string {
	if err := c.BodyParser(data); err != nil {
		paserQueryError := fmt.Sprintf("Unable to Parser Query : %v", err)
		return &paserQueryError
	}
	if err := StructValidator(data); err != nil {
		return err
	}
	return nil
}
