package core

import "time"

// Trip is an entity that represents launches for users
type Trip struct {
	Destination string
	Date        time.Time
}
