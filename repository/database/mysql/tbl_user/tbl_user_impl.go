package tbluser

import (
	"helle/entity/request"
	dbrepo "helle/repository/database"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) dbrepo.UserRepository {
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
