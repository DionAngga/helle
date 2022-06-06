package usecase

import (
	"helle/entity/database"
	repositorymysql "helle/repository/database"
)

type accUsecase struct {
	accRepository repositorymysql.UserAccountRepository
}

func New(accrepository repositorymysql.UserAccountRepository) *accUsecase {
	return &accUsecase{accrepository}
}

func (u *accUsecase) GetUsername(account string) (*database.TblUserAccount, error) {
	user, err := u.accRepository.FindUsername(account)
	if err != nil {
		return nil, err
	}
	return user, nil
}
