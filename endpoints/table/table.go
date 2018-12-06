package table

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/phucledien/cafe-pos/domain"
	"github.com/phucledien/cafe-pos/service"
)

// CreateData data for CreateTable
type CreateData struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// CreateRequest request struct for CreateTable
type CreateRequest struct {
	Table CreateData `json:"table"`
}

// CreateResponse response struct for CreateTable
type CreateResponse struct {
	Table domain.Table `json:"table"`
}

// StatusCode custom status code for success create Table
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Table
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req   = request.(CreateRequest)
			table = &domain.Table{
				Name:   req.Table.Name,
				Status: req.Table.Status,
			}
		)

		err := s.TableService.Create(ctx, table)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Table: *table}, nil
	}
}

// FindRequest request struct for Find a Table
type FindRequest struct {
	TableID domain.UUID
}

// FindResponse response struct for Find a Table
type FindResponse struct {
	Table *domain.Table `json:"table"`
}

// MakeFindEndpoint make endpoint for find table
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var tableFind domain.Table
		req := request.(FindRequest)
		tableFind.ID = req.TableID

		table, err := s.TableService.Find(ctx, &tableFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Table: table}, nil
	}
}

// FindAllRequest request struct for find all Table
type FindAllRequest struct{}

// FindAllResponse response struct for find all Table
type FindAllResponse struct {
	Tables []domain.Table `json:"tables"`
}

// MakeFindAllEndpoint make endpoint for find all Table
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		tables, err := s.TableService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Tables: tables}, nil
	}
}

// UpdateData data for Update Table
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	Name   string      `json:"name"`
	Status bool        `json:"status"`
}

// UpdateRequest request struct for update Table
type UpdateRequest struct {
	Table UpdateData `json:"table"`
}

// UpdateResponse response struct for update Table
type UpdateResponse struct {
	Table domain.Table `json:"table"`
}

// MakeUpdateEndpoint make endpoint for update a table
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req   = request.(UpdateRequest)
			table = domain.Table{
				Model:  domain.Model{ID: req.Table.ID},
				Name:   req.Table.Name,
				Status: req.Table.Status,
			}
		)

		res, err := s.TableService.Update(ctx, &table)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Table: *res}, nil
	}
}

// DeleteRequest request struct for delete a table
type DeleteRequest struct {
	TableID domain.UUID
}

// DeleteResponse response struct for delete a table
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for delete a Table
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       = request.(DeleteRequest)
			tableFind = domain.Table{}
		)
		tableFind.ID = req.TableID

		err := s.TableService.Delete(ctx, &tableFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
