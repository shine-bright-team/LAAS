package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
)

func RequestValidator(c *fiber.Ctx, data interface{}) *globalmodels.ErrorResponse {
	if err := c.BodyParser(data); err != nil {
		return &globalmodels.ErrorResponse{Type: "Parser", Message: fmt.Sprintf("Unable to Parser Query : %v", err)}
	}
	if err := StructValidator(data); err != nil {
		return err
	}
	return nil
}
