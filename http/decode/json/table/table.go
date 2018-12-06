package table

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/phucledien/cafe-pos/domain"
	tableEndpoint "github.com/phucledien/cafe-pos/endpoints/table"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	tableID, err := domain.UUIDFromString(chi.URLParam(r, "table_id"))
	if err != nil {
		return nil, err
	}
	return tableEndpoint.FindRequest{TableID: tableID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return tableEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req tableEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	tableID, err := domain.UUIDFromString(chi.URLParam(r, "table_id"))
	if err != nil {
		return nil, err
	}

	var req tableEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Table.ID = tableID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	tableID, err := domain.UUIDFromString(chi.URLParam(r, "table_id"))
	if err != nil {
		return nil, err
	}
	return tableEndpoint.DeleteRequest{TableID: tableID}, nil
}
