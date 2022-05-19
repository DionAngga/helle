package controller

import (
	"encoding/json"
	"helle/entity/request"
	"helle/usecase"
	"net/http"
)

type Controller interface {
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
	GetProfilebyUsername(w http.ResponseWriter, r *http.Request)
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
