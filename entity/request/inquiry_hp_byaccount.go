package request

type User struct {
	Client        string `json:"client" valid:"required"`
	AccountNumber string `json:"account_number"`
	RequestRefnum string `json:"request_refnum"`
}

type Name struct {
	Username string `json:"username" validate:"required"`
}
