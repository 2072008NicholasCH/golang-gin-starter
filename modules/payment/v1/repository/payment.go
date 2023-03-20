package repository

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/entity"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

type PaymentRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Payment, error)
	FindByID(ctx context.Context, id int64) (*entity.Payment, error)
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) FindAll(ctx context.Context) ([]*entity.Payment, error) {
	var payments []*entity.Payment

	err := r.db.Find(&payments).Error

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return payments, nil
}

func (r *PaymentRepository) FindByID(ctx context.Context, id int64) (*entity.Payment, error) {
	var payment *entity.Payment

	err := r.db.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return payment, nil
}
