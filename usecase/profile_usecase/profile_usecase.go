package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
	repositorymysql "helle/repository/database"
)

type profileUsecase struct {
	profileRepository repositorymysql.UserProfileRepository
}

func New(profilerespository repositorymysql.UserProfileRepository) *profileUsecase {
	return &profileUsecase{profilerespository}
}

func (u *profileUsecase) GetProfile(username *request.Name) (*database.TblUserProfile, error) {
	input := username
	user, err := u.profileRepository.FindProfile(input.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
