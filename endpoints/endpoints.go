package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phucledien/cafe-pos/service"

	"github.com/phucledien/cafe-pos/endpoints/drink"
	"github.com/phucledien/cafe-pos/endpoints/table"
	"github.com/phucledien/cafe-pos/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindTable    endpoint.Endpoint
	FindAllTable endpoint.Endpoint
	CreateTable  endpoint.Endpoint
	UpdateTable  endpoint.Endpoint
	DeleteTable  endpoint.Endpoint

	FindDrink    endpoint.Endpoint
	FindAllDrink endpoint.Endpoint
	CreateDrink  endpoint.Endpoint
	UpdateDrink  endpoint.Endpoint
	DeleteDrink  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindTable:    table.MakeFindEndpoint(s),
		FindAllTable: table.MakeFindAllEndpoint(s),
		CreateTable:  table.MakeCreateEndpoint(s),
		UpdateTable:  table.MakeUpdateEndpoint(s),
		DeleteTable:  table.MakeDeleteEndpoint(s),

		FindDrink:    drink.MakeFindEndpoint(s),
		FindAllDrink: drink.MakeFindAllEndpoint(s),
		CreateDrink:  drink.MakeCreateEndpoint(s),
		UpdateDrink:  drink.MakeUpdateEndpoint(s),
		DeleteDrink:  drink.MakeDeleteEndpoint(s),
	}
}
