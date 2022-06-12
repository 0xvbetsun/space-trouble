package psql

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/vbetsun/space-trouble/internal/core"
)

func TestUser_GetUser(t *testing.T) {
	db, mock := NewMock()
	repo := NewUser(db)
	defer func() {
		repo.db.Close()
	}()

	uuid := uuid.New().String()
	type args struct {
		firstName string
		lastName  string
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    core.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(uuid)
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("john", "doe").WillReturnRows(rows)
			},
			input: args{"john", "doe"},
			want: core.User{
				ID: uuid,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.GetUser(tt.input.firstName, tt.input.lastName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CreateUser(t *testing.T) {
	db, mock := NewMock()
	repo := NewUser(db)
	defer func() {
		repo.db.Close()
	}()

	uuid := uuid.New().String()
	birthday, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")
	if err != nil {
		t.Fatal("can't parse birthday")
	}
	tests := []struct {
		name    string
		mock    func()
		input   core.User
		want    core.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "gender", "birthday"}).AddRow(uuid, "Test", "Test", 0, birthday)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("Test", "Test", core.Male, birthday).WillReturnRows(rows)
			},
			input: core.User{
				FirstName: "Test",
				LastName:  "Test",
				Gender:    core.Male,
				Birthday:  birthday,
			},
			want: core.User{
				ID:        uuid,
				FirstName: "Test",
				LastName:  "Test",
				Gender:    core.Male,
				Birthday:  birthday,
			},
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "gender"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("", "", "", "").WillReturnRows(rows)
			},
			input: core.User{
				FirstName: "",
				LastName:  "",
				Gender:    0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.CreateUser(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
