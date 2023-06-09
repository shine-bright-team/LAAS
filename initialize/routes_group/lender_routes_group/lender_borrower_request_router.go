package lender_routes_group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
)

// /lender/borrower/request
func LenderBorrowerRequestRouter(router fiber.Router) {
	router.Get("/", lender_routes.GetBorrowerRequest)
	router.Patch("/", lender_routes.UpdateBorrowRequest)
	router.Post("/signature", lender_routes.UploadSignatureOnUpdateBorrowRequest)
	router.Get("/:borrowId", lender_routes.GetBorrowRequestByBorrowId)
}
