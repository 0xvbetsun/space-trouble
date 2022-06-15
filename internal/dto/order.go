// package dto is using for storing objects for sending data between layers in application
package dto

import "time"

// Order represents data-transfer object for order entity
type Order struct {
	UserID        string
	LaunchpadID   string
	DestinationID string
	LaunchDate    time.Time
}
