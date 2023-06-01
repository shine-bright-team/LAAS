package middlewere

import (
	"github.com/gofiber/fiber/v2"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"strconv"
	"strings"
)

func GetUserMiddleware(c *fiber.Ctx) error {
	errorResponse := c.Status(fiber.StatusUnauthorized).JSON(globalmodels.ErrorResponse{Type: "Unauthorized", Message: "Unauthorized, make sure you logged in"})
	authorizationHeader := c.Get("Authorization")

	splittedHeader := strings.Split(authorizationHeader, " ")
	if len(splittedHeader) != 2 && splittedHeader[0] != "Bearer" {
		return errorResponse
	}
	token := splittedHeader[1]
	if claims, error := utils.ValidateToken(token); error != nil {
		print(error.Error())
		return errorResponse
	} else {
		if s, err := strconv.Atoi(claims.UserId); err != nil {
			return errorResponse
		} else {
			c.Locals("userId", s)
			c.Locals("isLender", claims.IsLender)
		}
		return c.Next()
	}

}
