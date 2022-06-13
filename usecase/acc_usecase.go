package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
)

type AccUsecase interface {
	FindUsername(account *request.Acc, id *response.Response)
}
