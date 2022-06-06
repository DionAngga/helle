package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
)

type ProfileUseCase interface {
	GetProfile(username *request.Name) (*database.TblUserProfile, error)
}
