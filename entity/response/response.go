package response

type Response struct {
	ResponseCode   string      `json:"response_code"`
	ResponseDesc   string      `json:"response_desc"`
	ResponseId     string      `json:"response_id"`
	ResponseRefnum string      `json:"response_refnum"`
	ResponseData   interface{} `json:"response_data"`
}

type InquiryHp struct {
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

type Name struct {
	Username string `json:"username"`
}

type Validate struct {
	Validation string `json:"validation"`
	Field      string `json:"field"`
}

type Emtpy struct{}
