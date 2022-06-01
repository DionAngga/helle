package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
)

type Usecase interface {
	GetInquiry(client request.User) (*request.User, error)
	GetProfile(username *request.Name) (*database.TblUserProfile, error)
	GetUsername(account string) (*database.TblUserAccount, error)
	GetUserPhoneNumber(account *request.User) (*database.TblUserProfile, error)
}
