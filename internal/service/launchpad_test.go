package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/vbetsun/space-trouble/internal/core"
)

func TestLaunchpadService_GetForDate(t *testing.T) {
	type fields struct {
		storage LaunchpadStorage
	}
	type args struct {
		launchpadID string
		date        time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Launchpad
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LaunchpadService{
				storage: tt.fields.storage,
			}
			got, err := s.GetForDate(tt.args.launchpadID, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("LaunchpadService.GetForDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LaunchpadService.GetForDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
