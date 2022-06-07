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
	usecase usecase.UserUsecase
}

func New(usecase usecase.UserUsecase) *controller {
	return &controller{usecase}
}

func handleError(err interface{}) {
	if err != nil {
		fmt.Println("error")
	}
}

func (c *controller) GetInquirybyaccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.User
	err := json.NewDecoder(r.Body).Decode(&user)
	handleError(err)
	User, err := c.usecase.GetInquiry(user)
	handleError(err)
	err = json.NewEncoder(w).Encode(User)
	handleError(err)
}

func (c *controller) GetUserPhoneNumber(w http.ResponseWriter, r *http.Request) {
	user := request.User{}
	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	handleError(err)
	valid := validator.New()
	err = valid.Struct(user)
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	if err != nil {
		Response := response.Response{
			ResponseCode:   "VE",
			ResponseDesc:   "fail on validation",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: user.RequestRefnum,
			ResponseData:   err.Error(),
		}
		err = json.NewEncoder(w).Encode(Response)
		handleError(err)
		return

	}

	User, err := c.usecase.GetUserPhoneNumber(&user)
	if err != nil {
		Response := response.Response{
			ResponseCode:   "",
			ResponseDesc:   "error in gorm",
			ResponseId:     "",
			ResponseRefnum: user.RequestRefnum,
			ResponseData:   err,
		}
		err = json.NewEncoder(w).Encode(Response)
		handleError(err)
		return
	}
	if User.Username == "" {
		Response := response.Response{
			ResponseCode:   "AN",
			ResponseDesc:   "Account Number Not Found",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: user.RequestRefnum,
			ResponseData:   response.Emtpy{},
		}
		err = json.NewEncoder(w).Encode(Response)
		handleError(err)
		return
	}
	if User != nil {
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
		err = json.NewEncoder(w).Encode(&Response)
		handleError(err)
	}

}
