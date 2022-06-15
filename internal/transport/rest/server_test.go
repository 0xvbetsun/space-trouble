// Package rest implements setup, teardown and handlers for the REST API
package rest

import (
	"context"
	"net/http"
	"testing"
)

func TestServer_Run(t *testing.T) {
	type fields struct {
		httpServer *http.Server
	}
	type args struct {
		port    string
		handler http.Handler
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
			s := &Server{
				httpServer: tt.fields.httpServer,
			}
			if err := s.Run(tt.args.port, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Server.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Shutdown(t *testing.T) {
	type fields struct {
		httpServer *http.Server
	}
	type args struct {
		ctx context.Context
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
			s := &Server{
				httpServer: tt.fields.httpServer,
			}
			if err := s.Shutdown(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
