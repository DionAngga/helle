package tbluseraccount

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

func (r *repository) FindUsername(account string) (*database.TblUserAccount, error) {
	var acc database.TblUserAccount
	dbres := r.DB.Where("account = ?", account).First(&acc)
	return &acc, dbres.Error
}
