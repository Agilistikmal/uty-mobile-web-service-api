package repository

import (
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (s *PaymentRepository) Create(payment *model.Payment) (*model.Payment, error) {
	err := s.db.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentRepository) Update(id string, payment *model.Payment) (*model.Payment, error) {
	err := s.db.Where("id = ?", id).Updates(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentRepository) FindByID(id string) (*model.Payment, error) {
	var payment *model.Payment
	err := s.db.Take(&payment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentRepository) FindByReferenceID(referenceID string) (*model.Payment, error) {
	var payment *model.Payment
	err := s.db.Take(&payment, "reference_id = ?", referenceID).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}
