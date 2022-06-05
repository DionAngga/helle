package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
	repositorymysql "helle/repository/database"
)

type usecase struct {
	userRepository    repositorymysql.UserRepository
	accRepository     repositorymysql.UserAccountRepository
	profileRepository repositorymysql.UserProfileRepository
}

func New(repositorys repositorymysql.UserRepository, repositoryacc repositorymysql.UserAccountRepository, respositoryprofile repositorymysql.UserProfileRepository) *usecase {
	return &usecase{repositorys, repositoryacc, respositoryprofile}
}

func (u *usecase) GetInquiry(client request.User) (*request.User, error) {
	input := client
	user, err := u.userRepository.FindUser(input.Client)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usecase) GetProfile(username *request.Name) (*database.TblUserProfile, error) {
	input := username
	user, err := u.profileRepository.FindProfile(input.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usecase) GetUsername(account string) (*database.TblUserAccount, error) {
	user, err := u.accRepository.FindUsername(account)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usecase) GetUserPhoneNumber(account *request.User) (*database.TblUserProfile, error) {
	user, err := u.accRepository.FindUsername(account.AccountNumber)
	if err != nil {
		return nil, err
	}
	client, err := u.profileRepository.FindProfile(user.Username)
	if err != nil {
		return nil, err
	}
	return client, nil
}
