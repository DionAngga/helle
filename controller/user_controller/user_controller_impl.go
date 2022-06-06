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
)

type controller struct {
	usecase usecase.UserUsecase
}

func New(usecase usecase.UserUsecase) *controller {
	return &controller{usecase}
}

func (c *controller) GetInquirybyaccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.User
	json.NewDecoder(r.Body).Decode(&user)
	User, err := c.usecase.GetInquiry(user)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(User)
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
		Response := response.Response{
			ResponseCode:   "VE",
			ResponseDesc:   "fail on validation",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: user.RequestRefnum,
			ResponseData:   err.Error(),
		}
		json.NewEncoder(w).Encode(Response)
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
		json.NewEncoder(w).Encode(Response)
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
		json.NewEncoder(w).Encode(Response)
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
		json.NewEncoder(w).Encode(&Response)
	}

}
