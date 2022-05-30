package repositorymysql

import (
	"helle/entity/database"
	"helle/entity/request"

	"gorm.io/gorm"
)

type Repository interface {
	FindUser(client string) (*request.User, error)
	FindProfile(Username string) (*database.TblUserProfile, error)
	FindUsername(account string) (*database.TblUserAccount, error)
}

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *repository {
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

func (r *repository) FindProfile(username string) (*database.TblUserProfile, error) {
	var profile database.TblUserProfile
	err := r.DB.Where("username = ?", username).Find(&profile).Error
	if err != nil {
		return &profile, err
	}
	return &profile, nil
}

func (r *repository) FindUsername(account string) (*database.TblUserAccount, error) {
	var acc database.TblUserAccount
	err := r.DB.Where("account = ?", account).Find(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
