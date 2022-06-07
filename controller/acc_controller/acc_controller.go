package controller

import (
	"encoding/json"
	"fmt"
	"helle/entity/database"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
)

type controller struct {
	usecase usecase.AccUsecase
}

func New(usecase usecase.AccUsecase) *controller {
	return &controller{usecase}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *database.TblUserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer r.Body.Close()

	User, err := c.usecase.GetUsername(user.Account)
	if err != nil {
		respon := response.Response{}
		respon.ResponseCode = "99"
		_, _ = w.Write([]byte(respon.ResponseCode))
		return
	}

	// tambah error validation

	errs := json.NewEncoder(w).Encode(&User)
	if errs != nil {
		fmt.Println("error")
	}
}
