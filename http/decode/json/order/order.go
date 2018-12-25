package order

import (
	"context"
	"encoding/json"
	"net/http"

	// "github.com/go-chi/chi"

	// "github.com/phucledien/cafe-pos/domain"
	orderEndpoint "github.com/phucledien/cafe-pos/endpoints/order"
)

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return orderEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req orderEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
