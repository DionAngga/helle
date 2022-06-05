package tbluserprofile

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

func (r *repository) FindProfile(username string) (*database.TblUserProfile, error) {
	var profile database.TblUserProfile
	err := r.DB.Where("username = ?", username).Find(&profile).Error
	if err != nil {
		return &profile, err
	}
	return &profile, nil
}
