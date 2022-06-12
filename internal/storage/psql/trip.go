package psql

import "database/sql"

// Trip represents repository for trips entity
type Trip struct {
	db *sql.DB
}

// NewTrip return instance of trip repository
func NewTrip(db *sql.DB) *Trip {
	return &Trip{db}
}
