package service

import (
	"github.com/vbetsun/space-trouble/internal/core"
)

type UserStorage interface {
	CreateUser(core.User) (core.User, error)
	GetUser(username, password string) (core.User, error)
}

// UserService represents service for handling users in the system
type UserService struct {
	storage UserStorage
}

// NewUserService returns instance of service for handling users in the system
func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage}
}
