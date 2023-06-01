package globalmodels

type UserInfoResponse struct {
	Id        string `json:"id"`
	IsKyc     bool   `json:"is_kyc"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	isLender  bool   `json:"is_lender"`
}
