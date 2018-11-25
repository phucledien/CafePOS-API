package service

import "github.com/phucledien/cafe-pos/service/user"

// Service define list of all services in projects
type Service struct {
	UserService user.Service
}
