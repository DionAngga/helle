package response

import (
	"encoding/json"
	"fmt"
	loggers "helle/log"
	"net/http"
)

func init() {
	loggers.Init()
}

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

func New(id string) *Response {
	return &Response{
		ResponseCode:   "xx",
		ResponseDesc:   "general error",
		ResponseId:     id,
		ResponseRefnum: "",
		ResponseData:   Emtpy{},
	}
}

func (r *Response) SetResponseCode(code string) {
	r.ResponseCode = code
}

func (r *Response) SetResponseDesc(desc string) {
	r.ResponseDesc = desc
}

func (r *Response) SetResponseData(data interface{}) {
	r.ResponseData = data
}

func (r *Response) SetResponseRefnum(refnum string) {
	r.ResponseRefnum = refnum
}

func (r *Response) SendResponse(result *Response, w http.ResponseWriter) {
	var req *http.Request
	//var header http.Header
	//js, _ := json.Marshal(result)
	loggers.LogResponse(req, "ideapad e7", "Dion", "inquiry_hp_byaccount", "00", result, w.Header())
	fmt.Println("Headerrrrr==", w.Header())
	_ = json.NewEncoder(w).Encode(&result)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// return err
	// log := logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{
	// 	FieldMap: logrus.FieldMap{
	// 		logrus.FieldKeyTime: "timestamp",
	// 		logrus.FieldKeyMsg:  "message",
	// 	},
	// })
	// js, _ := json.Marshal(result)
	// log.Info("response: ", string(js))
}
