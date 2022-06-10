package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
	repositorymysql "helle/repository/database"
)

type accUsecase struct {
	accRepository repositorymysql.UserAccountRepository
}

func New(accrepository repositorymysql.UserAccountRepository) *accUsecase {
	return &accUsecase{accrepository}
}

func (u *accUsecase) FindUsername(rqst *request.Acc, id *response.Response) error {

	user, err := u.accRepository.FindUsername(rqst.AccountNumber)
	if err != nil || user == nil {
		Response := response.Response{
			ResponseCode:   "AN",
			ResponseDesc:   "Account Number Not Found",
			ResponseRefnum: rqst.RequestRefnum,
			ResponseData:   response.Emtpy{},
		}
		return err
	}

	Response := response.Response{
		ResponseCode:   "00",
		ResponseDesc:   "Get Phone By Accnum Success",
		ResponseRefnum: rqst.RequestRefnum,
		ResponseData:   response.Name{Username: user.Username},
	}
	return &Response, err

}
