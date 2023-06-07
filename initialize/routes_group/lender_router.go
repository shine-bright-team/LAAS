package routesgroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

func LenderRouter(router fiber.Router) {
	router.Post("/", lender_routes.CreateLenderPreference)

	router.Get("/borrower/request", lender_routes.GetBorrowerRequest)
	router.Get("/borrower", lender_routes.GetBorrowersDebt)

	router.Post("/borrower/request", lender_routes.DecideBorrowRequest)
	router.Get("/borrower/request/:borrowId", lender_routes.GetBorrowRequestByBorrowId)

	router.Get("/debt", lender_routes.GetDebt)
	router.Get("/debt/:debtId", lender_routes.GetDebtById)
	router.Post("/debt/transaction", lender_routes.DecideTransaction)
}
