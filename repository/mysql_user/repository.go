package repositorymysql

import (
	"fmt"
	"helle/entity/request"
	"helle/entity/response"

	"gorm.io/gorm"
)

type Repository interface {
	FindUser(client string) (*request.User, error)
	FindProfile(Username string) (*response.TblUserProfile, error)
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
	fmt.Println("inquiryy", inquiry)

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
