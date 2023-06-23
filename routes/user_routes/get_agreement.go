package userroutes

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/db"
	dbmodel "github.com/shine-bright-team/LAAS/v2/db/db_model"
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"github.com/shine-bright-team/LAAS/v2/routes/lender_routes"
	"gorm.io/gorm"
	"log"
)

func GetAgreemnt(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	var agreement dbmodel.Agreement
	if res := db.DB.Model(&agreement).Where("user_Id = ?", userId).First(&agreement); res.Error != nil {
		log.Print(res.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}
	var user dbmodel.User
	if res := db.DB.Model(&dbmodel.User{}).Where("id = ?", userId).First(user); res.Error != nil {
		log.Print(res.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	var InterestRateFormat *string
	var FormatAmountrange *string
	InterestRateFormat = nil
	if agreement.InterestRate != nil {
		tempFormat := fmt.Sprintf("%.2f%%", *agreement.InterestRate)
		if agreement.IsInterestPerMonth {
			tempFormat = tempFormat + " per month"
		} else {
			tempFormat = tempFormat + " per day"
		}
		InterestRateFormat = &tempFormat
	}
	if agreement.InterestRate != nil {
		amountrange := fmt.Sprintf("%.2f", agreement.LowestAmount)
		amountrange += " - " + fmt.Sprintf("%.2f", agreement.HighestAmount)
		FormatAmountrange = &amountrange
	}

	var aggregatedReview lender_routes.BorrowReviewAggregate

	if result := db.DB.Raw("select avg(score) as score, count(*) as review_count from reviews where reviewed_user_id = ?;", userId).Scan(&aggregatedReview); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			aggregatedReview = lender_routes.BorrowReviewAggregate{
				Score:       0,
				ReviewCount: 0,
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
		}
	}

	return c.JSON(globalmodels.AgreementResponse{
		UserId:       agreement.UserId,
		ID:           agreement.ID,
		AmountRange:  FormatAmountrange,
		InterestRate: InterestRateFormat,
		DueIn:        agreement.DueIn,
		Addition:     agreement.Addition,
		Review: globalmodels.ReviewResponse{
			ReviewAverage: aggregatedReview.Score,
			ReviewCount:   aggregatedReview.ReviewCount,
		},
		ActiveAtLeast:  agreement.ActiveAtLeast,
		HaveBaseSalary: agreement.BaseSalary,
		PaymentChannel: user.PayChannel,
		PaymentNumber:  user.PayNumber,
	})

}
