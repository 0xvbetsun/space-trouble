package service

type OrderStorage interface{}

// OrderService represents service for handling orders in the system
type OrderService struct {
	storage OrderStorage
}

// NewOrderService returns instance of service for handling orders in the system
func NewOrderService(storage OrderStorage) *OrderService {
	return &OrderService{storage}
}
