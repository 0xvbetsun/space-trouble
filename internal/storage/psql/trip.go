package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
)

// Trip represents repository for trips entity
type Trip struct {
	db *sql.DB
}

// NewTrip return instance of trip repository
func NewTrip(db *sql.DB) *Trip {
	return &Trip{db}
}

// GetByDestination returns trip entity mapped to the trips schedule
func (r *Trip) GetByDestination(destinationID string) (core.Trip, error) {
	var trip core.Trip
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, getTripQuery(), destinationID).Scan(&trip.Destination, &trip.Date)

	return trip, err
}

func getTripQuery() string {
	return fmt.Sprintf(`--sql
		SELECT d.name, ts.date
		FROM %s AS d
		LEFT JOIN %s AS ts ON d.id = ts.destination_id
		WHERE id = $1 
	`, destinationsTable, tripsScheduleTable)
}
