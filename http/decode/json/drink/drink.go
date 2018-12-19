package drink

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/phucledien/cafe-pos/domain"
	drinkEndpoint "github.com/phucledien/cafe-pos/endpoints/drink"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	drinkID, err := domain.UUIDFromString(chi.URLParam(r, "drink_id"))
	if err != nil {
		return nil, err
	}
	return drinkEndpoint.FindRequest{DrinkID: drinkID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return drinkEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req drinkEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	drinkID, err := domain.UUIDFromString(chi.URLParam(r, "drink_id"))
	if err != nil {
		return nil, err
	}

	var req drinkEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Drink.ID = drinkID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	drinkID, err := domain.UUIDFromString(chi.URLParam(r, "drink_id"))
	if err != nil {
		return nil, err
	}
	return drinkEndpoint.DeleteRequest{DrinkID: drinkID}, nil
}
