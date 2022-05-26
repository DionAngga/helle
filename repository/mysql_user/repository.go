package repositorymysql

import (
	"helle/entity/request"
	"helle/entity/response"

	"gorm.io/gorm"
)

type Repository interface {
	FindUser(client string) (*request.User, error)
	FindProfile(Username string) (*response.TblUserProfile, error)
	FindUsername(account string) (*response.TblUserAccount, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser(client string) (*request.User, error) {
	var inquiry request.User
	err := r.DB.Where("client = ?", client).Find(&inquiry).Error
	if err != nil {
		return &inquiry, err
	}

	return &inquiry, nil
}

func (r *repository) FindProfile(username string) (*response.TblUserProfile, error) {
	var profile response.TblUserProfile
	err := r.DB.Where("username = ?", username).Find(&profile).Error
	if err != nil {
		return &profile, err
	}
	return &profile, nil
}

func (r *repository) FindUsername(account string) (*response.TblUserAccount, error) {
	var acc response.TblUserAccount
	err := r.DB.Where("account = ?", account).Find(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
