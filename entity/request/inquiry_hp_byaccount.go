package request

type User struct {
	Client        string `json:"client" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required"`
}

type Name struct {
	Username string `json:"username" validate:"required"`
}
