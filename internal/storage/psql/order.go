package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

// Order represents repository for orders entity
type Order struct {
	db *sql.DB
}

// NewOrder return instance of order repository
func NewOrder(db *sql.DB) *Order {
	return &Order{db}
}

// GetOrderByID returns order by ID
func (r *Order) GetByID(orderID string) (core.Order, error) {
	var order core.Order
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.db.QueryRowContext(ctx, orderByIDQuery(), orderID).Scan(&order.ID)

	return order, err
}

// CreateOrder creates new order
func (r *Order) Create(userID string, data *dto.Order) (string, error) {
	var orderID string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx, createOrderQuery(), data.LaunchpadID, data.DestinationID, userID, data.LaunchDate,
	).Scan(&orderID)

	return orderID, err
}

// GetAllOrders returns all orders from DB
func (r *Order) GetAll() ([]core.Order, error) {
	var orders []core.Order
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, allOrdersQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order core.Order
		err := rows.Scan(
			&order.ID,
			&order.FirstName,
			&order.LastName,
			&order.Gender,
			&order.Birthday,
			&order.Launchpad,
			&order.Destination,
			&order.Status,
			&order.LaunchDate,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderByID returns order by ID
func (r *Order) RemoveByID(orderID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.db.ExecContext(ctx, removeOrderByIDQuery(), orderID)

	return err
}

func orderByIDQuery() string {
	return fmt.Sprintf(`--sql
		SELECT o.id
		FROM %s AS o 
		WHERE o.deleted_at IS NULL
		AND o.id = $1
	`, ordersTable)
}

func allOrdersQuery() string {
	return fmt.Sprintf(`--sql
		SELECT 
			o.id,
			u.first_name,
			u.last_name,
			u.gender,
			u.birthday,
			l.full_name AS launchpad, 
			d.name AS destination, 
			o.status,
			o.launch_date 
		FROM %s AS o 
		INNER JOIN %s AS u ON u.id = o.user_id 
		INNER JOIN %s AS d ON d.id = o.destination_id 
		INNER JOIN %s AS l ON l.id = o.launchpad_id 
		WHERE o.deleted_at IS NULL
	`, ordersTable, usersTable, destinationsTable, launchpadsTable)
}

func createOrderQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (launchpad_id, destination_id, user_id, launch_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, ordersTable)
}

func removeOrderByIDQuery() string {
	return fmt.Sprintf(`--sql
		UPDATE %s
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE o.deleted_at IS NULL
		AND o.id = $1
	`, ordersTable)
}
