// Package core represents domain's entities
package core

import "time"

type Gender int

const (
	Male Gender = iota
	Female
)

// User it is an entity of end user for this application
type User struct {
	ID        int       `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    Gender    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
}
