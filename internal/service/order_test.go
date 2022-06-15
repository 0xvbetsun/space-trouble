package service

import (
	"reflect"
	"testing"

	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

func TestOrderService_GetByID(t *testing.T) {
	type fields struct {
		storage OrderStorage
	}
	type args struct {
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderService{
				storage: tt.fields.storage,
			}
			got, err := s.GetByID(tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_GetAll(t *testing.T) {
	type fields struct {
		storage OrderStorage
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*core.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderService{
				storage: tt.fields.storage,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_Create(t *testing.T) {
	type fields struct {
		storage OrderStorage
	}
	type args struct {
		userID string
		data   *dto.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderService{
				storage: tt.fields.storage,
			}
			got, err := s.Create(tt.args.userID, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_RemoveByID(t *testing.T) {
	type fields struct {
		storage OrderStorage
	}
	type args struct {
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderService{
				storage: tt.fields.storage,
			}
			if err := s.RemoveByID(tt.args.orderID); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.RemoveByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
