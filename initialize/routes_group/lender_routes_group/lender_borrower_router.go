package lender_routes_group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

// /lender/borrower
func LenderBorrowerRouter(router fiber.Router) {
	router.Get("/", lender_routes.GetBorrowersDebt)

	borrowerRequestGroup := router.Group("/request")
	LenderBorrowerRequestRouter(borrowerRequestGroup)

}
