package service

type LaunchpadStorage interface{}

// LaunchpadService represents service for handling launchpads in the system
type LaunchpadService struct {
	storage LaunchpadStorage
}

// NewLaunchpadService returns instance of service for handling launchpad's data in the system
func NewLaunchpadService(storage LaunchpadStorage) *LaunchpadService {
	return &LaunchpadService{storage}
}
