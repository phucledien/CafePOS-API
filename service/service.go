package service

import (
	"github.com/phucledien/cafe-pos/service/drink"
	"github.com/phucledien/cafe-pos/service/order"
	"github.com/phucledien/cafe-pos/service/table"
	"github.com/phucledien/cafe-pos/service/user"
)

// Service define list of all services in projects
type Service struct {
	DrinkService drink.Service
	TableService table.Service
	UserService  user.Service
	OrderService order.Service
}
