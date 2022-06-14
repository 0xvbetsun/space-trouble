package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

// User represents repository for users entity
type User struct {
	db *sql.DB
}

// NewUser return instance of user repository
func NewUser(db *sql.DB) *User {
	return &User{db}
}

// Create creates new user in DB
func (r *User) Create(u *dto.User) (core.User, error) {
	var user core.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, createUserQuery(), u.FirstName, u.LastName, u.Gender, u.Birthday).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Birthday)

	return user, err
}

// GetByFullName returns user from DB by name
func (r *User) GetByFullName(firstName, lastName string) (core.User, error) {
	var user core.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, getUserQuery(), firstName, lastName).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Birthday)

	return user, err
}

func createUserQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (first_name, last_name, gender, birthday) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, first_name, last_name, gender, birthday
	`, usersTable)
}

func getUserQuery() string {
	return fmt.Sprintf(`--sql
		SELECT id, first_name, last_name, gender, birthday FROM %s 
		WHERE first_name = $1 
		AND last_name = $2
	`, usersTable)
}
