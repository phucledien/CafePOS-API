package service

import "github.com/phucledien/cafe-pos/service/user"
import "github.com/phucledien/cafe-pos/service/table"

// Service define list of all services in projects
type Service struct {
	UserService  user.Service
	TableService table.Service
}
