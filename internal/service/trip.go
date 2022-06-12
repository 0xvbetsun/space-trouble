package service

type TripStorage interface{}

// TripService represents service for handling trips in the system
type TripService struct {
	storage TripStorage
}

// NewTripService returns instance of service for handling trips' data in the system
func NewTripService(storage TripStorage) *TripService {
	return &TripService{storage}
}
