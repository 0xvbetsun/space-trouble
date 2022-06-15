package psql

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

func TestUser_FindOrCreate(t *testing.T) {
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
		input   dto.User
		want    core.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "gender", "birthday"}).AddRow(uuid, "Test", "Test", "male", birthday)
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("Test", "Test").WillReturnRows(rows)
			},
			input: dto.User{
				FirstName: "Test",
				LastName:  "Test",
				Gender:    "male",
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
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("", "", "", "").WillReturnRows(rows)
			},
			input: dto.User{
				FirstName: "",
				LastName:  "",
				Gender:    "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.FindOrCreate(&tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
