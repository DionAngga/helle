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

	id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(id, "-", "", -1)
	rspn := response.New(uuidWithoutHyphens)

	var rqst *request.Acc
	err := json.NewDecoder(r.Body).Decode(&rqst)
	if err != nil {
		rspn.SetResponseCode("GE")
		rspn.SetResponseDesc("General Error: " + err.Error())
		_ = json.NewEncoder(w).Encode(rspn)
		return
	}
	rqst.RequestId = uuidWithoutHyphens
	fmt.Println("rqst: ", rqst)
	valid := validator.New()
	err = valid.Struct(rqst)
	if err != nil {
		rspn.SetResponseCode("VE")
		rspn.SetResponseDesc("fail on validation")
		rspn.SetResponseData(err.Error())
		_ = json.NewEncoder(w).Encode(&rspn)
		return
	}

	rspn.SetResponseRefnum(rqst.RequestRefnum)
	err := c.usecase.FindUsername(rqst, rspn)
	if err != nil {
		fmt.Println("error")
	}
	_ = json.NewEncoder(w).Encode(&rspn)

}
