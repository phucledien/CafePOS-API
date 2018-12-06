package table

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/phucledien/cafe-pos/domain"
)

// pgService implement for Table service in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Table service
func (s *pgService) Create(_ context.Context, p *domain.Table) error {
	return s.db.Create(p).Error
}

// Update implement Update for Table service
func (s *pgService) Update(_ context.Context, p *domain.Table) (*domain.Table, error) {
	old := domain.Table{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Name = p.Name
	old.Status = p.Status

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Table service
func (s *pgService) Find(_ context.Context, p *domain.Table) (*domain.Table, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement Find All for Table service
func (s *pgService) FindAll(_ context.Context) ([]domain.Table, error) {
	res := []domain.Table{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Table service
func (s *pgService) Delete(_ context.Context, p *domain.Table) error {
	old := domain.Table{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	return s.db.Delete(old).Error
}
