package globalmodels

type UserInfoResponse struct {
	Id             uint   `json:"id"`
	IsKyc          bool   `json:"is_kyc"`
	Username       string `json:"username"`
	Title          string `json:"title"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	IsLender       bool   `json:"is_lender"`
	Token          string `json:"token"`
	IsSetAgreement bool   `json:"is_set_agreement"`
}
