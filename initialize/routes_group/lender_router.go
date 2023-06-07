package routesgroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/initialize/routes_group/lender_routes_group"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

func LenderRouter(router fiber.Router) {
	router.Post("/", lender_routes.CreateLenderPreference)
	borrowerGroup := router.Group("/borrower")
	debtGroup := router.Group("/debt")
	lender_routes_group.LenderBorrowerRouter(borrowerGroup)
	lender_routes_group.LenderDebtGroup(debtGroup)
}
