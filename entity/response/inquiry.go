package response

import "time"

type Inquiry struct {
	ResponseCode   string                 `json:"response_code"`
	ResponseDesc   string                 `json:"response_desc"`
	ResponseId     string                 `json:"response_id"`
	ResponseRefnum string                 `json:"response_refnum"`
	ResponseData   map[string]interface{} `json:"response_data"`
}

type TblUserProfile struct {
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	BornPlace        string    `json:"born_place"`
	BornDate         time.Time `json:"born_date"`
	MotherMaidenName string    `json:"mother_maiden_name"`
	Address          string    `json:"address"`
	CellphoneNumber  string    `json:"cellphone_number"`
	EmailAddress     string    `json:"email_address"`
	Cif              string    `json:"cif"`
}
