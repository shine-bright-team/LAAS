package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"time"
)

func GetBorrowersDebt(c *fiber.Ctx) error {
	a1 := 50.0
	timeN := time.Now().AddDate(0, 3, 0)
	user1 := globalmodels.BorrowRequestResponse{BorrowId: 5, UserId: 5, Firstname: "Muay", Lastname: "Mi", RequestedAt: time.Now(), RequestedAmount: 50, Username: "Mmuay", RemainingAmount: &a1, DueDate: &timeN}
	user2 := globalmodels.BorrowRequestResponse{BorrowId: 6, UserId: 6, Firstname: "Fa", Lastname: "H", RequestedAt: time.Now(), RequestedAmount: 300, Username: "Fxh", RemainingAmount: &a1, DueDate: &timeN}
	user3 := globalmodels.BorrowRequestResponse{BorrowId: 7, UserId: 7, Firstname: "Gu", Lastname: "Gun", RequestedAt: time.Now(), RequestedAmount: 3000, Username: "GuGun", RemainingAmount: &a1, DueDate: &timeN}
	return c.JSON([3]globalmodels.BorrowRequestResponse{user1, user2, user3})

}
