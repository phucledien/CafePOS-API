package drink

import (
	"context"
	"github.com/phucledien/cafe-pos/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Drink) error
	Update(ctx context.Context, p *domain.Drink) (*domain.Drink, error)
	Find(ctx context.Context, p *domain.Drink) (*domain.Drink, error)
	FindAll(ctx context.Context) ([]domain.Drink, error)
	Delete(ctx context.Context, p *domain.Drink) error
}
