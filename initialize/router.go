package initialize

import (
	"github.com/gofiber/fiber/v2"
	routesgroup "github.com/shine-bright-team/LAAS/v2/initialize/routes_group"
)

func Router(app *fiber.App) {
	user_group := app.Group("/user")
	auth_group := app.Group("/auth")
	routesgroup.UserRouter(user_group)
	routesgroup.AuthRouter(auth_group)

}
