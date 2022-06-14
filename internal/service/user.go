package service

import (
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

type UserStorage interface {
	Create(data *dto.User) (core.User, error)
	GetByFullName(firstName, lastName string) (core.User, error)
}

// UserService represents service for handling users in the system
type UserService struct {
	storage UserStorage
}

// NewUserService returns instance of service for handling users in the system
func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage}
}

func (s UserService) FindOrCreate(data *dto.User) (core.User, error) {
	var emptyUser core.User
	user, err := s.storage.GetByFullName(data.FirstName, data.LastName)
	if err != nil {
		return emptyUser, nil
	}
	if user != emptyUser {
		return user, err
	}
	return s.storage.Create(data)
}
