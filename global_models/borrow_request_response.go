package globalmodels

import "time"

type DebtAnalysisResponse struct {
	Paid   float64 `json:"paid"`
	UnPaid float64 `json:"unpaid"`
}

type ReviewResponse struct {
	ReviewAverage float64 `json:"review_average"`
	ReviewCount   int     `json:"review_count"`
}

type BorrowRequestResponse struct {
	BorrowId        uint                  `json:"borrow_id"`
	Username        string                `json:"username"`
	UserId          uint                  `json:"user_id"`
	Firstname       string                `json:"firstname"`
	Lastname        string                `json:"lastname"`
	RequestedAmount float64               `json:"requested_amount"`
	RemainingAmount *float64              `json:"remaining_amount"`
	RequestedAt     time.Time             `json:"requested_at"`
	DueDate         *time.Time            `json:"due_date"`
	PayChannel      *string               `json:"pay_channel"`
	PayNumber       *string               `json:"pay_number"`
	DebtAnalysis    *DebtAnalysisResponse `json:"debt_analysis"`
	Reviews         *ReviewResponse       `json:"reviews"`
}
