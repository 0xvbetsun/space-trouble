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

// FindOrCreate returns user from DB and creates it if user doesn't exist
func (r *User) FindOrCreate(data *dto.User) (core.User, error) {
	var user core.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx, findUserQuery(), data.FirstName, data.LastName,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Birthday)

	if err == sql.ErrNoRows {
		err = r.db.QueryRowContext(
			ctx, createUserQuery(), data.FirstName, data.LastName, data.Gender, data.Birthday,
		).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Birthday)
	}

	return user, err
}

func findUserQuery() string {
	return fmt.Sprintf(`--sql
		SELECT id, first_name, last_name, gender, birthday
		FROM %s
		WHERE first_name = $1
		AND last_name = $2
	`, usersTable)
}

func createUserQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (first_name, last_name, gender, birthday) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT ON CONSTRAINT users_first_name_last_name_key DO NOTHING
		RETURNING id, first_name, last_name, gender, birthday
	`, usersTable)
}
