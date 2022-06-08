package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
	repositorymysql "helle/repository/database"
	"strings"

	"github.com/google/uuid"
)

type profileUsecase struct {
	profileRepository repositorymysql.UserProfileRepository
	accRepository     repositorymysql.UserAccountRepository
}

func New(profilerespository repositorymysql.UserProfileRepository, accRepository repositorymysql.UserAccountRepository) *profileUsecase {
	return &profileUsecase{profilerespository, accRepository}
}

func (u *profileUsecase) FindProfile(rqst *request.Name) (*response.Response, error) {
	respon_id := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(respon_id, "-", "", -1)
	user, err := u.profileRepository.FindProfile(rqst.Username)
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
		ResponseData:   &user,
	}
	return &Response, nil
}

func (u *profileUsecase) FindUserPhoneNumber(rqst *request.User) (*response.Response, error) {
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

	rspn, err := u.profileRepository.FindProfile(user.Username)
	if err != nil || rspn.CellphoneNumber == "" {
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
		ResponseData:   &response.InquiryHp{PhoneNumber: rspn.CellphoneNumber, EmailAddress: rspn.EmailAddress},
	}
	return &Response, nil
}
