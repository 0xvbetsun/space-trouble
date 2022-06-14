// PAckage handler implements router and endpoints for REST API
package handler

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
	"go.uber.org/zap"
)

const layout = "2006-01-02"

type LaunchpadService interface {
	GetForDate(launchpadID string, date time.Time) (core.Launchpad, error)
}

type OrderService interface {
	GetByID(orderID string) (core.Order, error)
	GetAll() ([]*core.Order, error)
	Create(userID string, data *dto.Order) (core.Order, error)
	RemoveByID(orderID string) error
}

type TripService interface {
	GetByDestination(destinationID string) (core.Trip, error)
}

type UserService interface {
	FindOrCreate(data *dto.User) (core.User, error)
}

type Services struct {
	Launchpad LaunchpadService
	Order     OrderService
	Trip      TripService
	User      UserService
}

// Deps represents external dependencies for rest handlers
type Deps struct {
	Services
	Log *zap.Logger
}

// Handler represents rest modules of API
type Handler struct {
	Order *OrderHandler
	log   *zap.Logger
}

// New returns instance of rest handler
func New(deps Deps) *Handler {
	return &Handler{
		Order: NewOrderHandler(deps.Services, deps.Log),
		log:   deps.Log,
	}
}

// Routes creates, composes, and returns rest routes for API
func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(h.Logger())
	r.Use(h.Recoverer)
	r.Use(SendRequestID)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", h.Order.getAllOrders)
		r.Post("/", h.Order.createOrder)
		r.Route("/{orderID}", func(r chi.Router) {
			r.Use(h.Order.orderCtx)
			r.Delete("/", h.Order.deleteOrder)
		})
	})
	return r
}
