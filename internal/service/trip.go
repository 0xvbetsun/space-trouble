package service

import "github.com/vbetsun/space-trouble/internal/core"

type TripStorage interface {
	GetByDestination(destinationID string) (core.Trip, error)
}

// TripService represents service for handling trips in the system
type TripService struct {
	storage TripStorage
}

// NewTripService returns instance of service for handling trips' data in the system
func NewTripService(storage TripStorage) *TripService {
	return &TripService{storage}
}

// GetByDestination fetches trip by destination id
func (s *TripService) GetByDestination(destinationID string) (core.Trip, error) {
	return s.storage.GetByDestination(destinationID)
}
