package tbluser

import (
	"helle/entity/database"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser(client string) (*database.User, error) {
	var inquiry database.User
	err := r.DB.Where("client = ?", client).Find(&inquiry).Error
	if err != nil {
		return &inquiry, err
	}

	return &inquiry, nil
}
