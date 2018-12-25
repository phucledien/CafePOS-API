package order

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phucledien/cafe-pos/domain"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implmenter Create for Order service
func (s *pgService) Create(ctx context.Context, p *domain.Order) error {

	if err := s.db.Create(p).Error; err != nil {
		return err
	}

	if err := s.db.Preload("OrderDetails.Drink").Find(p).Error; err != nil {
		return err
	}

	totalPayment := 0
	for _, orderDetail := range p.OrderDetails {
		totalPayment += orderDetail.Drink.Price * orderDetail.Quantity
	}
	p.TotalPayment = totalPayment

	table := &domain.Table{Model: domain.Model{ID: p.TableID}}
	if err := s.db.Find(table).Error; err != nil {
		return err
	}
	table.Status = 1
	return s.db.Save(p).Save(table).Error
}

// Update(ctx context.Context, p *domain.Order) (*domain.Order, error)
// Find(ctx context.Context, p *domain.Order) (*domain.Order, error)

func (s *pgService) FindAll(ctx context.Context) ([]domain.Order, error) {
	res := []domain.Order{}
	return res, s.db.Preload("OrderDetails").Preload("OrderDetails.Drink").Find(&res).Error
}

// Delete(ctx context.Context, p *domain.Order) error
