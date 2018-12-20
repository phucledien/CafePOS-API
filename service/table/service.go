package table

import (
	"context"

	"github.com/phucledien/cafe-pos/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Table) error
	Update(ctx context.Context, p *domain.Table) (*domain.Table, error)
	UpdateStatus(ctx context.Context, tableID domain.UUID, status int) error
	Find(ctx context.Context, p *domain.Table) (*domain.Table, error)
	FindAll(ctx context.Context) ([]domain.Table, error)
	Delete(ctx context.Context, p *domain.Table) error
	GetEmptyTables(ctx context.Context) ([]domain.Table, error)
	GetPreparingTables(ctx context.Context) ([]domain.Table, error)
}
