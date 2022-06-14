// Package core represents domain's entities
package core

import "time"

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	}
	return "unknown"
}

// User it is an entity of end user for this application
type User struct {
	ID        string
	FirstName string
	LastName  string
	Gender    Gender
	Birthday  time.Time
	CreatedAt time.Time
}
