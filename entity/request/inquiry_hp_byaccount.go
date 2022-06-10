package request

type Acc struct {
	Client        string `json:"client" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required"`
	RequestId     string `json:"request_id" validate:"required"`
}

type Name struct {
	Client        string `json:"client" validate:"required"`
	Username      string `json:"username" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required"`
	RequestId     string `json:"request_id" validate:"required"`
}

type InquiryHpByAccount struct{}
