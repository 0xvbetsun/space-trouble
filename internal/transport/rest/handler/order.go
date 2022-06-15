package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/vbetsun/space-trouble/internal/core"
	"github.com/vbetsun/space-trouble/internal/dto"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Key to use when setting the order context.
type ctxKeyOrder string

const orderCtx ctxKeyOrder = "order"

// OrderHandler represents handler for rest endpoints related to order entity
type OrderHandler struct {
	services Services
	log      *zap.Logger
}
type CreateOrderRequest struct {
	*dto.Order
	*dto.User
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadID   string `json:"launchpad_id"`
	DestinationID string `json:"destination_id"`
	LaunchDate    string `json:"launch_date"`
}

type OrderResponse struct {
	*core.Order
}

// NewOrderHandler returns new instance of orders handler
func NewOrderHandler(services Services, log *zap.Logger) *OrderHandler {
	return &OrderHandler{services, log}
}

func (co *CreateOrderRequest) Bind(r *http.Request) error {
	// validation
	if co.FirstName == "" {
		return errors.New("first name is required")
	}
	if co.LastName == "" {
		return errors.New("last name is required")
	}
	if co.Gender == "" {
		return errors.New("gender is required")
	}
	if co.Birthday == "" {
		return errors.New("birthday is required")
	}
	if co.LaunchpadID == "" {
		return errors.New("id of launchpad is required")
	}
	if co.DestinationID == "" {
		return errors.New("id of destination is required")
	}
	if co.LaunchDate == "" {
		return errors.New("date of launch is required")
	}
	if co.Gender != core.Male && co.Gender != core.Female {
		return errors.New("gender should be 'male' or 'female'")
	}
	if _, err := uuid.Parse(co.DestinationID); err != nil {
		return errors.New("id of destination isn't correct")
	}
	if len(co.LaunchpadID) != 24 {
		return errors.New("id of launchpad isn't correct")
	}
	birthday, err := time.Parse(layout, co.Birthday)
	if err != nil {
		return fmt.Errorf("birthday should be in format %s", layout)
	}
	launchDate, err := time.Parse(layout, co.LaunchDate)
	if err != nil {
		return fmt.Errorf("date of launch should be in format %s", layout)
	}

	// normalization
	caser := cases.Title(language.English)
	firstName := caser.String(co.FirstName)
	lastName := caser.String(co.LastName)

	// population a DTOs for normalized data
	co.User.FirstName = firstName
	co.User.LastName = lastName
	co.User.Gender = co.Gender
	co.User.Birthday = birthday

	co.Order.LaunchpadID = co.LaunchpadID
	co.Order.DestinationID = co.DestinationID
	co.Order.LaunchDate = launchDate

	return nil
}

func (or *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func OrdersListResponse(orders []*core.Order) []render.Renderer {
	list := []render.Renderer{}
	for _, order := range orders {
		list = append(list, &OrderResponse{Order: order})
	}
	return list
}

func (h *OrderHandler) orderCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orderID, err := uuid.Parse(chi.URLParam(r, "orderID"))
		if err != nil {
			if rErr := render.Render(w, r, ErrRender(err)); rErr != nil {
				h.log.Error(rErr.Error())
			}
			return
		}
		order, err := h.services.Order.GetByID(orderID.String())
		if err != nil {
			if rErr := render.Render(w, r, ErrNotFound(ErrOrderNotFound)); rErr != nil {
				h.log.Error(rErr.Error())
			}
			return
		}
		ctx := context.WithValue(r.Context(), orderCtx, order)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *OrderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.services.Order.GetAll()
	if err != nil {
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	if err := render.RenderList(w, r, OrdersListResponse(orders)); err != nil {
		h.log.Error(err.Error())
	}
}

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	data := &CreateOrderRequest{Order: &dto.Order{}, User: &dto.User{}}
	if err := render.Bind(r, data); err != nil {
		h.log.Error(err.Error())
		if rErr := render.Render(w, r, ErrInvalidRequest(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	trip, err := h.services.Trip.GetByDestination(data.DestinationID)
	if err != nil {
		h.log.Error(err.Error())
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	if trip == (core.Trip{}) {
		if rErr := render.Render(w, r, ErrNotFound(ErrTripNotFound)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	if trip.Date.IsZero() {
		errNoSchedule := errors.New("sorry, but we haven't already had scheduled the requested date")
		if rErr := render.Render(w, r, ErrNotFound(errNoSchedule)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	launchpad, err := h.services.Launchpad.GetForDate(data.LaunchpadID, data.Order.LaunchDate)
	if err != nil {
		h.log.Error(err.Error())
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	if launchpad == (core.Launchpad{}) {
		if rErr := render.Render(w, r, ErrNotFound(ErrLaunchpadNotFound)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	if launchpad.Reserved {
		errReserved := fmt.Errorf("sorry, but launchpad %s is not available on that day", launchpad.ID)
		if rErr := render.Render(w, r, ErrNotFound(errReserved)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	user, err := h.services.User.FindOrCreate(data.User)
	if err != nil {
		h.log.Error(err.Error())
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	order, err := h.services.Order.Create(user.ID, data.Order)
	if err != nil {
		h.log.Error(err.Error())
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, &OrderResponse{Order: &order}); err != nil {
		h.log.Error(err.Error())
	}
}

func (h *OrderHandler) deleteOrder(w http.ResponseWriter, r *http.Request) {
	order, ok := r.Context().Value(orderCtx).(core.Order)
	if !ok {
		if err := render.Render(w, r, ErrInternalServer(ErrOrderCtxEmpty)); err != nil {
			h.log.Error(err.Error())
		}
		return
	}
	err := h.services.Order.RemoveByID(order.ID)
	if err != nil {
		if rErr := render.Render(w, r, ErrInternalServer(err)); rErr != nil {
			h.log.Error(rErr.Error())
		}
		return
	}
	render.NoContent(w, r)
}
