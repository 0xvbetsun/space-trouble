package psql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/vbetsun/space-trouble/internal/core"
)

func TestTrip_GetByDestination(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		destinationID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Trip
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Trip{
				db: tt.fields.db,
			}
			got, err := r.GetByDestination(tt.args.destinationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trip.GetByDestination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trip.GetByDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
