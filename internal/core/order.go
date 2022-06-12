// Package core represents domain's entities
package core

import "time"

type Status int

const (
	Active Status = iota
	Cancelled
	Finished
)

// Order it is an entity that represents user's booking of space trip
type Order struct {
	ID            int       `json:"-"`
	LaunchpadID   string    `json:"launchpad_id"`
	DestinationID string    `json:"destination_id"`
	Status        Status    `json:"status"`
	LaunchDate    time.Time `json:"launch_date"`
}
