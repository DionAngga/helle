package database

import (
	"helle/entity/database"
)

type UserProfileRepository interface {
	FindProfile(username string) (*database.TblUserProfile, error)
}
