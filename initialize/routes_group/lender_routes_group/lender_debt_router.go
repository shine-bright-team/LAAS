package lender_routes_group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

// /lender/debt
func LenderDebtGroup(router fiber.Router) {
	router.Get("/:debtId", lender_routes.GetDebtById)
	router.Get("/:debtId/:transactionId/image", lender_routes.GetTransactionPic)
	router.Patch("/transaction", lender_routes.UpdateTransaction)
}
