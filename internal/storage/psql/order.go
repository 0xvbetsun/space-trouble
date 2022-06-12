package psql

import (
	"database/sql"
)

// Order represents repository for orders entity
type Order struct {
	db *sql.DB
}

// NewOrder return instance of order repository
func NewOrder(db *sql.DB) *Order {
	return &Order{db}
}
