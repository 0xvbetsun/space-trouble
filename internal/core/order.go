// Package core represents domain's entities
package core

import "time"

type Status int

const (
	Active Status = iota
	Cancelled
)

// Order it is an entity that represents user's booking of space trip
type Order struct {
	ID          string    `json:"id"`
	Launchpad   string    `json:"launchpad"`
	Destination string    `json:"destination"`
	Status      Status    `json:"status"`
	LaunchDate  time.Time `json:"launch_date"`
	DeletedAt   time.Time `json:"-"`
}
