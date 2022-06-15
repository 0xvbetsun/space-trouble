package psql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
)

func TestOrder_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			r := &Order{
				db: tt.fields.db,
			}
			got, err := r.GetByID(tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Order.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		userID string
		data   *dto.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Order{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.userID, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Order.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Order.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []core.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Order{
				db: tt.fields.db,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Order.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_RemoveByID(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			r := &Order{
				db: tt.fields.db,
			}
			if err := r.RemoveByID(tt.args.orderID); (err != nil) != tt.wantErr {
				t.Errorf("Order.RemoveByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
