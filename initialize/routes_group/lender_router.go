package routesgroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

func LenderRouter(router fiber.Router) {
	router.Get("/borrowers/requests", lender_routes.GetBorrowerRequest)
	router.Get("/borrowers", lender_routes.GetBorrowersDebt)
}
