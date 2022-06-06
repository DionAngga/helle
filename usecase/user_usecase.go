package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
)

type UserUsecase interface {
	GetInquiry(client *request.User) (*database.User, error)
	GetUserPhoneNumber(account *request.User) (*database.TblUserProfile, error)
}
