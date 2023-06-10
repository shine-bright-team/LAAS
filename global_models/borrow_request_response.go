package globalmodels

import "time"

type BorrowRequestResponse struct {
	BorrowId        uint       `json:"borrow_id"`
	Username        string     `json:"username"`
	UserId          uint       `json:"user_id"`
	Firstname       string     `json:"firstname"`
	Lastname        string     `json:"lastname"`
	RequestedAmount int64      `json:"requested_amount"`
	RemainingAmount *int64     `json:"remaining_amount"`
	RequestedAt     time.Time  `json:"requested_at"`
	DueDate         *time.Time `json:"due_date"`
}
