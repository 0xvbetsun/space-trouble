package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
)

// Launchpad represents repository for launchpad entity
type Launchpad struct {
	db *sql.DB
}

// NewLaunchpad return instance of launchpad repository
func NewLaunchpad(db *sql.DB) *Launchpad {
	return &Launchpad{db}
}

// GetForDate fetches data from DB about launchpad for requested date
func (r *Launchpad) GetForDate(launchpadID string, date time.Time) (core.Launchpad, error) {
	var launchpad core.Launchpad
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, getLaunchpadQuery(), launchpadID, date).
		Scan(&launchpad.ID, &launchpad.Name, &launchpad.FullName, &launchpad.Reserved)

	return launchpad, err
}

// AddReservation adds reservation of requested launchpad to db
func (r *Launchpad) AddReservation(launchpadID, date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, addReservationQuery(), launchpadID, date)

	return err
}

func getLaunchpadQuery() string {
	return fmt.Sprintf(`--sql
		SELECT 
			l.id, 
			l.name, 
			l.full_name, 
			CASE WHEN lr.date IS NOT NULL
				THEN true
				ELSE false
			END reserved
		FROM %s AS l
		LEFT JOIN %s AS lr ON l.id = lr.launchpad_id AND lr.date = $2
		WHERE l.id = $1 
	`, launchpadsTable, launchpadReservationTable)
}

func addReservationQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (launchpad_id, date) 
		VALUES ($1, $2) 
		ON CONFLICT DO NOTHING
	`, launchpadReservationTable)
}
