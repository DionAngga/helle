package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
)

type ProfileUseCase interface {
	FindProfile(username *request.Name) (*response.Response, error)
	FindUserPhoneNumber(account *request.User) (*response.Response, error)
}
