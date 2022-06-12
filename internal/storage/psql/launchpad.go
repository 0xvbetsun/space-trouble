package psql

import "database/sql"

// Launchpad represents repository for launchpad entity
type Launchpad struct {
	db *sql.DB
}

// NewLaunchpad return instance of launchpad repository
func NewLaunchpad(db *sql.DB) *Launchpad {
	return &Launchpad{db}
}
