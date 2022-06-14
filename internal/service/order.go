package service

import (
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

type OrderStorage interface {
	GetByID(orderID string) (core.Order, error)
	GetAll() ([]core.Order, error)
	Create(userID string, data *dto.Order) (string, error)
	RemoveByID(orderID string) error
}

// OrderService represents service for handling orders in the system
type OrderService struct {
	storage OrderStorage
}

// NewOrderService returns instance of service for handling orders in the system
func NewOrderService(storage OrderStorage) *OrderService {
	return &OrderService{storage}
}

// GetByID handles business logic of fetching single order by its ID
func (s *OrderService) GetByID(orderID string) (core.Order, error) {
	return s.storage.GetByID(orderID)
}

// GetAll handles business logic of fetching all not deleted orders in the system
func (s *OrderService) GetAll() ([]core.Order, error) {
	return s.storage.GetAll()
}

// Create handles business logic of creating new order
func (s *OrderService) Create(userID string, data *dto.Order) (core.Order, error) {
	orderID, err := s.storage.Create(userID, data)
	if err != nil {
		return core.Order{}, err
	}
	return s.storage.GetByID(orderID)
}

// RemoveByID handles business logic of removing single order by its ID
func (s *OrderService) RemoveByID(orderID string) error {
	return s.storage.RemoveByID(orderID)
}
