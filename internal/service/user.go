package service

import (
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

type UserStorage interface {
	FindOrCreate(data *dto.User) (core.User, error)
}

// UserService represents service for handling users in the system
type UserService struct {
	storage UserStorage
}

// NewUserService returns instance of service for handling users in the system
func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage}
}

// FindOrCreate returns existed user by first & last name or creates new one
func (s UserService) FindOrCreate(data *dto.User) (core.User, error) {
	return s.storage.FindOrCreate(data)
}
