package order

import (
	"context"

	"github.com/phucledien/cafe-pos/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Order) error
	// Update(ctx context.Context, p *domain.Order) (*domain.Order, error)
	// Find(ctx context.Context, p *domain.Order) (*domain.Order, error)
	FindAll(ctx context.Context) ([]domain.Order, error)
	// Delete(ctx context.Context, p *domain.Order) error
}
