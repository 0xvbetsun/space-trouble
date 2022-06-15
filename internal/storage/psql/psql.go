// Package psql used for storing data in PostgreSQL database
package psql

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/log/zapadapter"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

const (
	usersTable                = "users"
	ordersTable               = "orders"
	launchpadsTable           = "launchpads"
	destinationsTable         = "destinations"
	tripsScheduleTable        = "trips_schedule"
	launchpadReservationTable = "launchpad_reservations"
)

// Config represents all required fields for connecting to postgres db
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Logger   *zap.Logger
}

// Storage contains all implemented repositories
type Storage struct {
	Launchpad *Launchpad
	Order     *Order
	Trip      *Trip
	User      *User
}

// String returns connection string from config
func (cfg Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
	)
}

// NewDB creates and validates new db connection
func NewDB(cfg Config) (*sql.DB, error) {
	zapadapter.NewLogger(cfg.Logger)
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// NewStorage returns all implemented repositories
func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Launchpad: NewLaunchpad(db),
		Order:     NewOrder(db),
		Trip:      NewTrip(db),
		User:      NewUser(db),
	}
}
