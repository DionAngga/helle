package controller

import (
	"encoding/json"
	"fmt"
	"helle/entity/request"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Controller interface {
	PostUser(w http.ResponseWriter, r *http.Request)
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
	GetProfilebyUsername(w http.ResponseWriter, r *http.Request)
	GetUsernameByAccount(w http.ResponseWriter, r *http.Request)
	GetUserPhoneNumber(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) Controller {
	return &controller{usecase}
}

func (c *controller) PostUser(w http.ResponseWriter, r *http.Request) {
	user := request.User{}
	json.NewDecoder(r.Body).Decode(&user)
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		Response := response.Inquiry{
			ResponseCode: "01",
			ResponseDesc: "Invalid Request",
			ResponseId:   "",
			ResponseData: response.Validate{Validation: "required", Field: "username"},
		}

		json.NewEncoder(w).Encode(Response)

	}
}

func (c *controller) GetInquirybyaccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user request.User
	json.NewDecoder(r.Body).Decode(&user)
	User, err := c.usecase.GetInquiry(user)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(User)
}

func (c *controller) GetProfilebyUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.Name

	json.NewDecoder(r.Body).Decode(&user)
	validate := validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		w.Header().Add("Content-Type", "application/json")
		responseBody := response.Validate{Validation: "required", Field: "username"}
		json.NewEncoder(w).Encode(responseBody)

	} else {
		User, err := c.usecase.GetProfile(user)
		if err != nil {
			responseBody := response.Validate{Validation: "error", Field: "gorm"}
			json.NewEncoder(w).Encode(responseBody)
		}
		if User.Username == "" {
			responseBody := response.Validate{Validation: "error", Field: "username not found"}
			json.NewEncoder(w).Encode(responseBody)
		} else {
			json.NewEncoder(w).Encode(&User)
		}
	}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *response.TblUserAccount
	json.NewDecoder(r.Body).Decode(&user)

	User, err := c.usecase.GetUsername(user.Account)
	fmt.Println("user===", User)
	if err != nil {
		return
	}

	resp := &response.Name{
		Username: User.Username,
	}

	json.NewEncoder(w).Encode(resp)
}

func (c *controller) GetUserPhoneNumber(w http.ResponseWriter, r *http.Request) {
	user := request.User{}
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	valid := validator.New()
	err := valid.Struct(user)
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	if err != nil {
		if user == (request.User{}) {
			Response := response.Inquiry{
				ResponseCode:   "AN",
				ResponseDesc:   "required validation failed on request",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "all request"},
			}
			json.NewEncoder(w).Encode(Response)
		} else if user.Client == "" {
			Response := response.Inquiry{
				ResponseCode:   "VE",
				ResponseDesc:   "required validation failed on client",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "client"},
			}
			json.NewEncoder(w).Encode(Response)
		} else if user.AccountNumber == "" {
			Response := response.Inquiry{
				ResponseCode:   "VE",
				ResponseDesc:   "required validation failed on account_number",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "account_number"},
			}
			json.NewEncoder(w).Encode(Response)
		} else if user.RequestRefnum == "" {
			Response := response.Inquiry{
				ResponseCode:   "VE",
				ResponseDesc:   "required validation failed on request_refnum",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "request_refnum"},
			}
			json.NewEncoder(w).Encode(Response)
		}

	} else {
		User, err := c.usecase.GetUserPhoneNumber(&user)
		if err != nil {
			Response := response.Inquiry{
				ResponseCode:   "00",
				ResponseDesc:   "error di gorm",
				ResponseId:     "",
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "username"},
			}
			json.NewEncoder(w).Encode(Response)
		}
		if User.Username == "" {
			Response := response.Inquiry{
				ResponseCode:   "AN",
				ResponseDesc:   "Account Number Not Found",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Emtpy{},
			}

			json.NewEncoder(w).Encode(Response)
		} else {
			Response := response.Inquiry{
				ResponseCode:   "00",
				ResponseDesc:   "Get Phone By Accnum Success",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData: response.InquiryHp{
					PhoneNumber:  User.CellphoneNumber,
					EmailAddress: User.EmailAddress,
				},
			}
			json.NewEncoder(w).Encode(&Response)
		}
	}
}
