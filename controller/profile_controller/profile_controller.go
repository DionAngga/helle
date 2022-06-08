package profilecontroller

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
	usecase usecase.ProfileUseCase
}

func New(usecase usecase.ProfileUseCase) *controller {
	return &controller{usecase}
}

func (c *controller) GetProfilebyUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.Name
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error")
	}
	valid := validator.New()
	err = valid.Struct(user)
	if err != nil {
		Response := response.Response{
			ResponseCode:   "VE",
			ResponseDesc:   "fail on validation",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: user.RequestRefnum,
			ResponseData:   err.Error(),
		}
		_ = json.NewEncoder(w).Encode(Response)
		return
	}
	User, err := c.usecase.FindProfile(user)
	if err != nil {
		fmt.Println("error")
	}

	_ = json.NewEncoder(w).Encode(&User)
}

func (c *controller) GetUserPhoneNumber(w http.ResponseWriter, r *http.Request) {
	rqst := request.User{}
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&rqst)
	valid := validator.New()
	err := valid.Struct(rqst)
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	if err != nil {
		Response := response.Response{
			ResponseCode:   "VE",
			ResponseDesc:   "fail on validation",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: rqst.RequestRefnum,
			ResponseData:   err.Error(),
		}
		_ = json.NewEncoder(w).Encode(Response)
		return

	}

	user, err := c.usecase.FindUserPhoneNumber(&rqst)
	if err != nil {
		_ = json.NewEncoder(w).Encode(user)
		return
	}
	_ = json.NewEncoder(w).Encode(&user)
}
