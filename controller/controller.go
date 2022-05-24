package controller

import (
	"encoding/json"
	"fmt"
	"helle/entity/request"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
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

func NewController(usecase usecase.Usecase) Controller {
	return &controller{usecase}
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
	User, err := c.usecase.GetProfile(user)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(&User)
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

	var user *response.TblUserAccount
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("user===", user.Account)
	User, err := c.usecase.GetUserPhoneNumber(user.Account)
	// fmt.Println("user===", reqnum)
	if err != nil {
		return
	}

	ResponData := &response.InquiryHp{
		PhoneNumber:  User.CellphoneNumber,
		EmailAddress: User.EmailAddress,
	}
	ResponDatas := &response.InquiryHp{
		PhoneNumber:  User.CellphoneNumber,
		EmailAddress: User.EmailAddress,
	}

	if ResponData != nil {
		Response := &response.Inquiry{
			ResponseCode:   "00",
			ResponseDesc:   "Get Phone By Accnum Success",
			ResponseRefnum: User.ResponseId,
			ResponseData:   *ResponData,
		}

		json.NewEncoder(w).Encode(Response)
	} else {
		Response := &response.Inquiry{
			ResponseCode: "01",
			ResponseData: *ResponDatas,
		}
		json.NewEncoder(w).Encode(Response)
	}

}
