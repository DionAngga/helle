package controller

import (
	"encoding/json"
	"helle/entity/request"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type controller struct {
	usecase usecase.AccUsecase
}

func New(usecase usecase.AccUsecase) *controller {
	return &controller{usecase}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(id, "-", "", -1)
	rspn := response.New(uuidWithoutHyphens)
	var rqst *request.Acc
	log.Info("request: ", rqst)
	err := json.NewDecoder(r.Body).Decode(&rqst)
	if err != nil {
		rspn.SetResponseCode("GE")
		rspn.SetResponseDesc("General Error: " + err.Error())
		rspn.SendResponse(w)
		return
	}
	rqst.RequestId = uuidWithoutHyphens

	valid := validator.New()
	err = valid.Struct(rqst)
	if err != nil {
		rspn.SetResponseCode("VE")
		rspn.SetResponseDesc("fail on validation")
		rspn.SetResponseData(err.Error())
		_ = json.NewEncoder(w).Encode(&rspn)
		log.Error("response: fail on validation")
		return
	}

	rspn.SetResponseRefnum(rqst.RequestRefnum)
	c.usecase.FindUsername(rqst, rspn)
	_ = json.NewEncoder(w).Encode(&rspn)
	log.Info("response: ", rspn)

}
