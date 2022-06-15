package service

import (
	"reflect"
	"testing"

	"github.com/vbetsun/space-trouble/internal/core"
)

func TestTripService_GetByDestination(t *testing.T) {
	type fields struct {
		storage TripStorage
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
			s := &TripService{
				storage: tt.fields.storage,
			}
			got, err := s.GetByDestination(tt.args.destinationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TripService.GetByDestination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TripService.GetByDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
