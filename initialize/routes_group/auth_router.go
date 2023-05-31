package routesgroup

import (
	"github.com/gofiber/fiber/v2"
	authroutes "github.com/shine-bright-team/LAAS/v2/routes/auth_routes"
)

func AuthRouter(router fiber.Router) {
	router.Post("/login", authroutes.Login)
	router.Post("/register", authroutes.Register)
}
