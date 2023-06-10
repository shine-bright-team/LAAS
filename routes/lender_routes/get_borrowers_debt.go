package lender_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
)

type contractWithRemaining struct {
	dbmodel.Contract
	RemainingAmount int
}

func GetBorrowersDebt(c *fiber.Ctx) error {

	userId := c.Locals("userId").(int)

	var contracts []contractWithRemaining

	if result := db.DB.Raw("SELECT *, loan_amount - (select sum(transactions.paid_amount) from transactions where contract_id = contracts.id) as remaining_amount from contracts join users u on u.id = contracts.borrower_user_id where lender_user_id = ?;", userId).Scan(&contracts); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	var responses []globalmodels.BorrowRequestResponse

	for i := range contracts {
		responses[i] = globalmodels.BorrowRequestResponse{
			BorrowId:        contracts[i].ID,
			Username:        contracts[i].Borrower.Username,
			UserId:          contracts[i].Borrower.ID,
			Firstname:       contracts[i].Borrower.Firstname,
			Lastname:        contracts[i].Borrower.Lastname,
			RequestedAmount: contracts[i].LoanAmount,
			RemainingAmount: &contracts[i].RemainingAmount,
			RequestedAt:     contracts[i].CreatedAt,
			DueDate:         &contracts[i].DueAt,
		}
	}

	return c.JSON(responses)

	//a1 := 50.0
	//timeN := time.Now().AddDate(0, 3, 0)
	//user1 := globalmodels.BorrowRequestResponse{BorrowId: 5, UserId: 5, Firstname: "Muay", Lastname: "Mi", RequestedAt: time.Now(), RequestedAmount: 50, Username: "Mmuay", RemainingAmount: &a1, DueDate: &timeN}
	//user2 := globalmodels.BorrowRequestResponse{BorrowId: 6, UserId: 6, Firstname: "Fa", Lastname: "H", RequestedAt: time.Now(), RequestedAmount: 300, Username: "Fxh", RemainingAmount: &a1, DueDate: &timeN}
	//user3 := globalmodels.BorrowRequestResponse{BorrowId: 7, UserId: 7, Firstname: "Gu", Lastname: "Gun", RequestedAt: time.Now(), RequestedAmount: 3000, Username: "GuGun", RemainingAmount: &a1, DueDate: &timeN}
	//return c.JSON([3]globalmodels.BorrowRequestResponse{user1, user2, user3})

}
