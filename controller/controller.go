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

//var users = []request.Name{}

func (c *controller) PostUser(w http.ResponseWriter, r *http.Request) {
	user := request.Name{}
	json.NewDecoder(r.Body).Decode(&user)
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		//validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(http.StatusBadRequest)
		// bukan ditulis validation error
		responseBody := response.Validate{Validation: "required", Field: "username"}
		json.NewEncoder(w).Encode(responseBody)

	}
	// We don't want an API user to set the ID manually
	// in a production use case this could be an automatically
	// ID in the database
	//users = append(users, user)
	// w.WriteHeader(http.StatusCreated)
}

//var validate *validator.Validate

func (c *controller) GetInquirybyaccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//validate = validator.New()
	var user request.User
	json.NewDecoder(r.Body).Decode(&user)
	User, err := c.usecase.GetInquiry(user)
	if err != nil {
		return
	}
	//fmt.Println("client", User)
	json.NewEncoder(w).Encode(User)
}

func (c *controller) GetProfilebyUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//validate = validator.New()
	var user *request.Name
	json.NewDecoder(r.Body).Decode(&user)
	validate := validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		//validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(http.StatusBadRequest)
		// bukan ditulis validation error
		responseBody := response.Validate{Validation: "required", Field: "username"}
		json.NewEncoder(w).Encode(responseBody)

	} else {
		User, err := c.usecase.GetProfile(user)
		if err != nil {
			responseBody := response.Validate{Validation: "required", Field: "username"}
			json.NewEncoder(w).Encode(responseBody)
		}
		if User.Username == "" {
			responseBody := response.Validate{Validation: "required", Field: "username"}
			json.NewEncoder(w).Encode(responseBody)
		} else {
			json.NewEncoder(w).Encode(&User)
		}
	}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//validate = validator.New()

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
	w.Header().Set("Content-Type", "application/json")
	//validate = validator.New()
	var usero *request.User

	json.NewDecoder(r.Body).Decode(&usero)
	validate := validator.New()
	errs := validate.Struct(usero)
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	if errs != nil {
		Response := response.Inquiry{
			ResponseCode:   "AN",
			ResponseDesc:   "salah ngga isi parameter",
			ResponseId:     "",
			ResponseRefnum: "",
			ResponseData:   response.Validate{Validation: "required", Field: "username"},
		}

		json.NewEncoder(w).Encode(Response)

	} else {
		User, err := c.usecase.GetUserPhoneNumber(usero)
		if err != nil {
			Response := response.Inquiry{
				ResponseCode:   "00",
				ResponseDesc:   "error di gorm",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: usero.RequestRefnum,
				ResponseData:   response.Validate{Validation: "required", Field: "username"},
			}
			json.NewEncoder(w).Encode(Response)
		}
		if User.Username == "" {
			Response := response.Inquiry{
				ResponseCode:   "AN",
				ResponseDesc:   "error salah isi",
				ResponseId:     "",
				ResponseRefnum: "",
				ResponseData:   response.Validate{},
			}

			json.NewEncoder(w).Encode(Response)
		} else {
			Response := response.Inquiry{
				ResponseCode:   "00",
				ResponseDesc:   "Get Phone By Accnum Success",
				ResponseId:     uuidWithoutHyphens,
				ResponseRefnum: usero.RequestRefnum,
				ResponseData: response.InquiryHp{
					PhoneNumber:  User.CellphoneNumber,
					EmailAddress: User.EmailAddress,
				},
			}
			json.NewEncoder(w).Encode(&Response)
		}
	}
}
