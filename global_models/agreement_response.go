package globalmodels

type AgreementResponse struct {
	UserId         uint
	ID             uint
	AmountRange    *string
	InterestRate   *string
	DueIn          int32
	Addition       string
	Review         ReviewResponse `json:"review"`
	ActiveAtLeast  *string        `json:"active_at_least"`
	HaveBaseSalary *string        `json:"have_base_salary"`
	PaymentChannel string         `json:"payment_channel"`
	PaymentNumber  string         `json:"payment_number"`
}
