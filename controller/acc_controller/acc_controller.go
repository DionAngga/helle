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

type controller struct {
	usecase usecase.AccUsecase
}

func New(usecase usecase.AccUsecase) *controller {
	return &controller{usecase}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rqst *request.User
	err := json.NewDecoder(r.Body).Decode(&rqst)

	if err != nil {
		fmt.Println("error")
	}

	valid := validator.New()
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	rspn := response.Response{}
	rspn.New(respon_id, rqst.RequestRefnum)

	err = valid.Struct(rqst)
	if err != nil {
		Response := response.Response{
			ResponseCode:   "VE",
			ResponseDesc:   "fail on validation",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: rqst.RequestRefnum,
			ResponseData:   err.Error(),
		}
		_ = json.NewEncoder(w).Encode(&Response)
		return
	}

	user, err := c.usecase.FindUsername(rqst)
	if err != nil {
		fmt.Println("error")
	}
	_ = json.NewEncoder(w).Encode(&user)

}
