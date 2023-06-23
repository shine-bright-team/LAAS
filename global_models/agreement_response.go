package globalmodels

type AgreementResponse struct {
	UserId         uint
	ID             uint
	AmountRange    *string
	InterestRate   *string
	DueIn          int32
	Addition       string
	Review         ReviewResponse `json:"review"`
	ActiveAtLeast  *int16         `json:"active_at_least"`
	HaveBaseSalary *int32         `json:"have_base_salary"`
	PaymentChannel string         `json:"payment_channel"`
	PaymentNumber  string         `json:"payment_number"`
}
