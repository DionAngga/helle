package controller

import (
	"encoding/json"
	"helle/entity/request"
	"helle/usecase"
	"net/http"
)

type controller struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) *controller {
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
