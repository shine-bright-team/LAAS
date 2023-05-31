package routesgroup

import "github.com/gofiber/fiber/v2"

func UserRouter(router fiber.Router) {
	router.Get("/")
	router.Post("/kyc")
}
