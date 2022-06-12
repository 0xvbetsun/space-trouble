package service

type Deps struct {
	LaunchpadStorage LaunchpadStorage
	OrderStorage     OrderStorage
	TripStorage      TripStorage
	UserStorage      UserStorage
}

// Service represents struct for handling business logic of the application
type Service struct {
	Launchpad *LaunchpadService
	Order     *OrderService
	Trip      *TripService
	User      *UserService
}

// NewService returns instance of service
func NewService(deps Deps) *Service {
	return &Service{
		Launchpad: NewLaunchpadService(deps.LaunchpadStorage),
		Order:     NewOrderService(deps.OrderStorage),
		Trip:      NewTripService(deps.TripStorage),
		User:      NewUserService(deps.UserStorage),
	}
}
