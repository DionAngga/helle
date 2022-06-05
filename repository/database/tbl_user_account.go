package database

import (
	"helle/entity/database"
)

type UserAccountRepository interface {
	FindUsername(account string) (*database.TblUserAccount, error)
}
