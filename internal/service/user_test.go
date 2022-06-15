package service

import (
	"reflect"
	"testing"

	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

func TestUserService_FindOrCreate(t *testing.T) {
	type fields struct {
		storage UserStorage
	}
	type args struct {
		data *dto.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				storage: tt.fields.storage,
			}
			got, err := s.FindOrCreate(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.FindOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FindOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
