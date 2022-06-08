package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
)

type AccUsecase interface {
	FindUsername(account *request.User) (*response.Response, error)
}
