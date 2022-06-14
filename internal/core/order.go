// Package core represents domain's entities
package core

import "time"

type Status int

const (
	Active Status = iota
	Canceled
)

func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Canceled:
		return "canceled"
	}
	return "unknown"
}

// Order it is an entity that represents user's booking of space trip
type Order struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      Gender    `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	Launchpad   string    `json:"launchpad"`
	Destination string    `json:"destination"`
	Status      Status    `json:"status"`
	LaunchDate  time.Time `json:"launch_date"`
	DeletedAt   time.Time `json:"-"`
}
