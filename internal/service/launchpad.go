package service

import (
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
)

type LaunchpadResponse struct {
	Launches []string `json:"launches"`
}

type LaunchpadStorage interface {
	GetForDate(launchpadID string, date time.Time) (core.Launchpad, error)
	AddReservation(launchpadID, date string) error
}

// LaunchpadService represents service for handling launchpads in the system
type LaunchpadService struct {
	storage LaunchpadStorage
}

// NewLaunchpadService returns instance of service for handling launchpad's data in the system
func NewLaunchpadService(storage LaunchpadStorage) *LaunchpadService {
	return &LaunchpadService{storage}
}

// GetForDate fetches launchpad information by ID for requested date
func (s *LaunchpadService) GetForDate(launchpadID string, date time.Time) (core.Launchpad, error) {
	return s.storage.GetForDate(launchpadID, date)
}
