package service

type Deps struct {
	UserStorage  UserStorage
	OrderStorage OrderStorage
}

type Service struct {
	User  *UserService
	Order *OrderService
}

func NewService(deps Deps) *Service {
	return &Service{
		User:  NewUserService(deps.UserStorage),
		Order: NewOrderService(deps.OrderStorage),
	}
}
