package lender_routes_group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

// /lender/debt
func LenderDebtGroup(router fiber.Router) {
	router.Get("/", lender_routes.GetDebt)
	router.Get("/:debtId", lender_routes.GetDebtById)
	router.Post("/transaction", lender_routes.DecideTransaction)
}