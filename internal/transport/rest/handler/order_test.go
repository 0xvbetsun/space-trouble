package handler

import (
	"net/http"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestOrderHandler_orderCtx(t *testing.T) {
	type fields struct {
		services Services
		log      *zap.Logger
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &OrderHandler{
				services: tt.fields.services,
				log:      tt.fields.log,
			}
			if got := h.orderCtx(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderHandler.orderCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderHandler_getAllOrders(t *testing.T) {
	type fields struct {
		services Services
		log      *zap.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &OrderHandler{
				services: tt.fields.services,
				log:      tt.fields.log,
			}
			h.getAllOrders(tt.args.w, tt.args.r)
		})
	}
}

func TestOrderHandler_createOrder(t *testing.T) {
	type fields struct {
		services Services
		log      *zap.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &OrderHandler{
				services: tt.fields.services,
				log:      tt.fields.log,
			}
			h.createOrder(tt.args.w, tt.args.r)
		})
	}
}

func TestOrderHandler_deleteOrder(t *testing.T) {
	type fields struct {
		services Services
		log      *zap.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &OrderHandler{
				services: tt.fields.services,
				log:      tt.fields.log,
			}
			h.deleteOrder(tt.args.w, tt.args.r)
		})
	}
}
