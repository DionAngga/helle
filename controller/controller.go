package controller

import (
	"encoding/json"
	"helle/entity/database"
	"helle/entity/request"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Controller interface {
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
	GetProfilebyUsername(w http.ResponseWriter, r *http.Request)
	GetUsernameByAccount(w http.ResponseWriter, r *http.Request)
	GetUserPhoneNumber(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) Controller {
	return &controller{usecase}
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
		responseBody := errs
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

	var user *database.TblUserAccount
	json.NewDecoder(r.Body).Decode(&user)

	User, err := c.usecase.GetUsername(user.Account)
	if err != nil {
		responseBody := response.Validate{Validation: "error", Field: "gorm"}
		json.NewEncoder(w).Encode(responseBody)
	}

	if User.Username == "" {
		responseBody := response.Validate{Validation: "error", Field: "username not found"}
		json.NewEncoder(w).Encode(responseBody)
	} else {
		json.NewEncoder(w).Encode(User)
	}
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
		if user.Client == "" || user.AccountNumber == "" || user.RequestRefnum == "" {
			Response := response.Response{
				ResponseCode:   "VE",
				ResponseDesc:   "fail on validation",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   err.Error(),
			}
			json.NewEncoder(w).Encode(Response)
		}
	} else {
		User, err := c.usecase.GetUserPhoneNumber(&user)
		if err != nil {
			Response := response.Response{
				ResponseCode:   "",
				ResponseDesc:   "error in gorm",
				ResponseId:     "",
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   err,
			}
			json.NewEncoder(w).Encode(Response)
		}
		if User.Username == "" {
			Response := response.Response{
				ResponseCode:   "AN",
				ResponseDesc:   "Account Number Not Found",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: user.RequestRefnum,
				ResponseData:   response.Emtpy{},
			}

			json.NewEncoder(w).Encode(Response)
		} else {
			Response := response.Response{
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
