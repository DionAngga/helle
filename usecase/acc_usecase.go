package usecase

import (
	"helle/entity/database"
)

type AccUsecase interface {
	GetUsername(account string) (*database.TblUserAccount, error)
}
