package drink

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/phucledien/cafe-pos/domain"
	"github.com/phucledien/cafe-pos/service"
)

// CreateData data for CreateDrink
type CreateData struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// CreateRequest request struct for CreateDrink
type CreateRequest struct {
	Drink CreateData `json:"drink"`
}

// CreateResponse response struct for CreateDrink
type CreateResponse struct {
	Drink domain.Drink `json:"drink"`
}

// StatusCode custom status code for success create drink
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Drink
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req   = request.(CreateRequest)
			drink = &domain.Drink{
				Name:  req.Drink.Name,
				Price: req.Drink.Price,
			}
		)

		err := s.DrinkService.Create(ctx, drink)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Drink: *drink}, nil
	}
}

// FindRequest request struct for Find a Drink
type FindRequest struct {
	DrinkID domain.UUID
}

// FindResponse response struct for Find a Drink
type FindResponse struct {
	Drink *domain.Drink `json:"drink"`
}

// MakeFindEndpoint make endpoint for find drink
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindRequest)
		findDrink := domain.Drink{Model: domain.Model{ID: req.DrinkID}}

		drink, err := s.DrinkService.Find(ctx, &findDrink)
		if err != nil {
			return nil, err
		}

		return FindResponse{Drink: drink}, nil
	}
}

// FindAllRequest request struct for find all Drink
type FindAllRequest struct{}

// FindAllResponse response struct for find all Drink
type FindAllResponse struct {
	Drinks []domain.Drink `json:"drinks"`
}

// MakeFindAllEndpoint make endpoint for find all Drink
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		drinks, err := s.DrinkService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Drinks: drinks}, nil
	}
}

// UpdateData data for Update Drink
type UpdateData struct {
	ID    domain.UUID `json:"-"`
	Name  string      `json:"name"`
	Price int         `json:"price"`
}

// UpdateRequest request struct for update Drink
type UpdateRequest struct {
	Drink UpdateData `json:"drink"`
}

// UpdateResponse response struct for update Drink
type UpdateResponse struct {
	Drink domain.Drink `json:"drink"`
}

// MakeUpdateEndpoint make endpoint for update a drink
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req   = request.(UpdateRequest)
			drink = domain.Drink{
				Model: domain.Model{ID: req.Drink.ID},
				Name:  req.Drink.Name,
				Price: req.Drink.Price,
			}
		)

		res, err := s.DrinkService.Update(ctx, &drink)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Drink: *res}, nil
	}
}

// DeleteRequest request struct for delete a drink
type DeleteRequest struct {
	DrinkID domain.UUID
}

// DeleteResponse response struct for delete a drink
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for delete a Drink
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       = request.(DeleteRequest)
			drinkFind = domain.Drink{}
		)
		drinkFind.ID = req.DrinkID

		err := s.DrinkService.Delete(ctx, &drinkFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
