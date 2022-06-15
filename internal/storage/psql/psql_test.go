package psql

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestConfig_String(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
		Logger   *zap.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Correct connection string",
			fields: fields{"localhost", "5432", "postgres", "pass123", "postgres", "disable", nil},
			want:   "host=localhost port=5432 user=postgres dbname=postgres password=pass123 sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DBName:   tt.fields.DBName,
				SSLMode:  tt.fields.SSLMode,
				Logger:   tt.fields.Logger,
			}
			if got := cfg.String(); got != tt.want {
				t.Errorf("Config.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
