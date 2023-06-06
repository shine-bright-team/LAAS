package initialize

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/initialize/middlewere"
	routesgroup "github.com/shine-bright-team/LAAS/v2/initialize/routes_group"
)

func Router(app *fiber.App) {
	user_group := app.Group("/user")
	auth_group := app.Group("/auth")
	lender_group := app.Group("/lender")
	routesgroup.AuthRouter(auth_group)
	app.Use(middlewere.GetUserMiddleware)
	routesgroup.UserRouter(user_group)
	app.Use(middlewere.VerifyLender)
	routesgroup.LenderRouter(lender_group)
}
