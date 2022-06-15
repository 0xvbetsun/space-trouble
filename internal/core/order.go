// Package core represents domain's entities
package core

import "time"

const (
	Active   = "active"
	Canceled = "canceled"
)

// Order it is an entity that represents user's booking of space trip
type Order struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	Launchpad   string    `json:"launchpad"`
	Destination string    `json:"destination"`
	Status      string    `json:"status"`
	LaunchDate  time.Time `json:"launch_date"`
	DeletedAt   time.Time `json:"-"`
}
