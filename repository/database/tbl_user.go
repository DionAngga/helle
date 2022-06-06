package database

import (
	"helle/entity/database"
)

type UserRepository interface {
	FindUser(client string) (*database.User, error)
}
