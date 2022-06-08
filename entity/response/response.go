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

type User struct {
	Client          string `json:"client"`
	AccountNumber   string `json:"account_number"`
	Timestamp       string `json:"timestamp"`
	ResponseRefnum  string `json:"response_refnum"`
	ResponseId      string `json:"response_id"`
	CellphoneNumber string `json:"cellphone_number"`
	EmailAddress    string `json:"email_address"`
	Refnum          string `json:"refnum"`
}

type Validate struct {
	Validation string `json:"validation"`
	Field      string `json:"field"`
}

type Emtpy struct{}

func (r *Response) New(id, refnum string) *Response {
	return &Response{
		ResponseCode:   "xx",
		ResponseDesc:   "general error",
		ResponseId:     id,
		ResponseRefnum: refnum,
		ResponseData:   Emtpy{},
	}
}
