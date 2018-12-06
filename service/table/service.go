package table

import (
	"context"

	"github.com/phucledien/cafe-pos/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Table) error
	Update(ctx context.Context, p *domain.Table) (*domain.Table, error)
	Find(ctx context.Context, p *domain.Table) (*domain.Table, error)
	FindAll(ctx context.Context) ([]domain.Table, error)
	Delete(ctx context.Context, p *domain.Table) error
}
