package order

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/phucledien/cafe-pos/domain"
	"github.com/phucledien/cafe-pos/service"
)

// CreateData data for CreateOrder
type CreateData struct {
	TableID      domain.UUID          `json:"table_id"`
	OrderDetails []domain.OrderDetail `json:"order_details"`
}

// CreateRequest request struct for CreateOrder
type CreateRequest struct {
	Order CreateData `json:"order"`
}

// CreateResponse response struct for CreateOrder
type CreateResponse struct {
	Order domain.Order `json:"order"`
}

// StatusCode customstatus code for success create User
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Order
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req   = request.(CreateRequest)
			order = &domain.Order{
				TableID:      req.Order.TableID,
				OrderDetails: req.Order.OrderDetails,
			}
		)

		err := s.OrderService.Create(ctx, order)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Order: *order}, nil
	}
}

// FindAllRequest request struct for FindAll Order
type FindAllRequest struct{}

// FindAllResponse request struct for find all Order
type FindAllResponse struct {
	Orders []domain.Order `json:"orders"`
}

// MakeFindAllEndpoint make endpoint for find all User
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		orders, err := s.OrderService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Orders: orders}, nil
	}
}
