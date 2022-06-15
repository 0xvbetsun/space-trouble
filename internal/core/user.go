// Package core represents domain's entities
package core

import (
	"time"
)

const (
	Male   = "male"
	Female = "female"
)

// User it is an entity of end user for this application
type User struct {
	ID        string
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
	CreatedAt time.Time
}
