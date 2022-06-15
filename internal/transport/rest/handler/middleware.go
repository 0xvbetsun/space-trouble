package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

// Logger is a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return.
func (h *Handler) Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				h.log.Info("Served",
					zap.String("proto", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("reqId", middleware.GetReqID(r.Context())))
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Recoverer is a middleware for recovering after panics during handling http requests
func (h *Handler) Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}
				if err := render.Render(w, r, ErrInternalServer(errors.New("something went wrong"))); err != nil {
					h.log.Error(err.Error())
				}
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// SendRequestID is a middleware for sending X-Request-Id to the client
func SendRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(
			middleware.RequestIDHeader,
			middleware.GetReqID(r.Context()),
		)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
