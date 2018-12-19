package drink

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phucledien/cafe-pos/domain"
)

// pgService implmenter for Drink serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Drink service
func (s *pgService) Create(ctx context.Context, p *domain.Drink) error {
	return s.db.Create(p).Error
}

// Update implement Update for Drink service
func (s *pgService) Update(ctx context.Context, p *domain.Drink) (*domain.Drink, error) {
	old := domain.Drink{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	old.Name = p.Name
	old.Price = p.Price

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Drink service
func (s *pgService) Find(ctx context.Context, p *domain.Drink) (*domain.Drink, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

// FindAll implement Find All for Drink service
func (s *pgService) FindAll(ctx context.Context) ([]domain.Drink, error) {
	res := []domain.Drink{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Drink service
func (s *pgService) Delete(ctx context.Context, p *domain.Drink) error {
	old := domain.Drink{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
