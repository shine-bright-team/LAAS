package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"time"
)

func GetBorrowerRequest(c *fiber.Ctx) error {
	user1 := globalmodels.BorrowRequestResponse{BorrowId: 1, UserId: 1, Firstname: "Sittichok", Lastname: "Ouamsiri", RequestedAt: time.Now(), RequestedAmount: 50, Username: "Thistine", RemainingAmount: nil, DueDate: nil}
	user2 := globalmodels.BorrowRequestResponse{BorrowId: 2, UserId: 3, Firstname: "Sittichok1", Lastname: "Ouamsiri", RequestedAt: time.Now(), RequestedAmount: 300, Username: "Thistine1", RemainingAmount: nil, DueDate: nil}
	user3 := globalmodels.BorrowRequestResponse{BorrowId: 2, UserId: 3, Firstname: "Sittichok2", Lastname: "Ouamsiri", RequestedAt: time.Now(), RequestedAmount: 3000, Username: "Thistine2", RemainingAmount: nil, DueDate: nil}

	return c.JSON([3]globalmodels.BorrowRequestResponse{user1, user2, user3})
}
