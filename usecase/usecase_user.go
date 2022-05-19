package usecase

import (
	"helle/entity/request"
	"helle/entity/response"
	repo "helle/repository/mysql_user"
)

type Usecase interface {
	GetInquiry(client request.User) (*request.User, error)
	GetProfile(username *request.Name) (*response.TblUserProfile, error)
}

type usercase struct {
	repository repo.Repository
}

func NewUser(repository repo.Repository) *usercase {
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

func (u *usercase) GetProfile(username *request.Name) (*response.TblUserProfile, error) {
	input := username
	user, err := u.repository.FindProfile(input.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
