package usecase

import (
	"helle/entity/database"
	"helle/entity/request"
	repo "helle/repository/database"
)

type Usecase interface {
	GetInquiry(client request.User) (*request.User, error)
	GetProfile(username *request.Name) (*database.TblUserProfile, error)
	GetUsername(account string) (*database.TblUserAccount, error)
	GetUserPhoneNumber(account *request.User) (*database.TblUserProfile, error)
}

type usercase struct {
	repository repo.Repository
}

func New(repository repo.Repository) *usercase {
	return &usercase{repository}
}

func (u *usercase) GetInquiry(client request.User) (*request.User, error) {
	input := client
	user, err := u.repository.FindUser(input.Client)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usercase) GetProfile(username *request.Name) (*database.TblUserProfile, error) {
	input := username
	user, err := u.repository.FindProfile(input.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usercase) GetUsername(account string) (*database.TblUserAccount, error) {
	user, err := u.repository.FindUsername(account)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usercase) GetUserPhoneNumber(account *request.User) (*database.TblUserProfile, error) {
	user, err := u.repository.FindUsername(account.AccountNumber)
	if err != nil {
		return nil, err
	}
	client, err := u.repository.FindProfile(user.Username)
	if err != nil {
		return nil, err
	}
	return client, nil
}
