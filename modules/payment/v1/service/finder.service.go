package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/payment/v1/repository"
)

type PaymentFinder struct {
	pfg         config.Config
	paymentRepo repository.PaymentRepositoryUseCase
}

type PaymentFinderUseCase interface {
	GetPayments(ctx context.Context) ([]*entity.Payment, error)
	GetPaymentByID(ctx context.Context, id int64) (*entity.Payment, error)
}

func NewPaymentFinder(pfg config.Config, paymentRepo repository.PaymentRepositoryUseCase) *PaymentFinder {
	return &PaymentFinder{
		pfg:         pfg,
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentFinder) GetPayments(ctx context.Context) ([]*entity.Payment, error) {
	payments, err := s.paymentRepo.FindAll(ctx)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return payments, nil
}

func (s *PaymentFinder) GetPaymentByID(ctx context.Context, id int64) (*entity.Payment, error) {
	payment, err := s.paymentRepo.FindByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return payment, nil
}
