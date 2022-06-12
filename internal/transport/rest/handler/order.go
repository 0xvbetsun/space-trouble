package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/service"
	"go.uber.org/zap"
)

// Key to use when setting the order context.
type ctxKeyOrder string

const orderCtx ctxKeyOrder = "order"

// OrderHandler represents handler for rest endpoints related to order entity
type OrderHandler struct {
	services *service.Service
	log      *zap.Logger
}

// NewOrderHandler returns new instance of orders handler
func NewOrderHandler(services *service.Service, log *zap.Logger) *OrderHandler {
	return &OrderHandler{services, log}
}

func (h *OrderHandler) orderCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), orderCtx, nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *OrderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *OrderHandler) deleteOrder(w http.ResponseWriter, r *http.Request) {
	order, ok := r.Context().Value(orderCtx).(core.Order)
	fmt.Println(order, ok)
	render.NoContent(w, r)
}
