package routesgroup

import (
	"github.com/gofiber/fiber/v2"
	userroutes "github.com/shine-bright-team/LAAS/v2/routes/user_routes"
)

func UserRouter(router fiber.Router) {
	router.Get("/", userroutes.GetCurrentUser)
	router.Post("/kyc", userroutes.UploadKyc)
}
