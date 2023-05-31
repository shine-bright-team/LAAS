package routesgroup

import "github.com/gofiber/fiber/v2"

func AuthRouter(router fiber.Router) {
	router.Post("/login")
	router.Post("/register")
}
