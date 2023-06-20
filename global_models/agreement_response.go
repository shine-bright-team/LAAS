package globalmodels

type AgreementResponse struct {
	UserId       uint
	ID           uint
	AmountRange  *string
	InterestRate *string
	DueIn        int32
	Addition     string
}
