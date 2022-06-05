package database

import (
	"helle/entity/request"
)

type UserRepository interface {
	FindUser(client string) (*request.User, error)
}
