package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
	repositorymysql "helle/repository/database"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type accUsecase struct {
	accRepository repositorymysql.UserAccountRepository
}

func New(accrepository repositorymysql.UserAccountRepository) *accUsecase {
	return &accUsecase{accrepository}
}

func (u *accUsecase) FindUsername(rqst *request.Acc, rspn *response.Response) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	user, err := u.accRepository.FindUsername(rqst.AccountNumber)

	if err != nil && err == gorm.ErrRecordNotFound {
		rspn.SetResponseCode("AN")
		rspn.SetResponseDesc("Account Number not found")
		return
	}

	if err != nil {
		rspn.SetResponseCode("DF")
		rspn.SetResponseDesc("Database Failure: " + err.Error())
		return
	}

	rspn.SetResponseCode("00")
	rspn.SetResponseDesc("Get Phone By Accnum Success")
	rspn.SetResponseData(response.Name{Username: user.Username})

}
