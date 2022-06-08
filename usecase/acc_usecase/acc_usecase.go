package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
	repositorymysql "helle/repository/database"
	"strings"

	"github.com/google/uuid"
)

type accUsecase struct {
	accRepository repositorymysql.UserAccountRepository
}

func New(accrepository repositorymysql.UserAccountRepository) *accUsecase {
	return &accUsecase{accrepository}
}

func (u *accUsecase) FindUsername(rqst *request.User) (*response.Response, error) {
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)

	user, err := u.accRepository.FindUsername(rqst.AccountNumber)
	if err != nil || user == nil {
		Response := response.Response{
			ResponseCode:   "AN",
			ResponseDesc:   "Account Number Not Found",
			ResponseId:     uuidWithoutHyphens,
			ResponseRefnum: rqst.RequestRefnum,
			ResponseData:   response.Emtpy{},
		}
		return &Response, err
	}

	Response := response.Response{
		ResponseCode:   "00",
		ResponseDesc:   "Get Phone By Accnum Success",
		ResponseId:     uuidWithoutHyphens,
		ResponseRefnum: rqst.RequestRefnum,
		ResponseData:   response.Name{Username: user.Username},
	}
	return &Response, err

}
