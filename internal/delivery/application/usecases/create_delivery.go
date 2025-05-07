package usecases

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type CreateDeliveryUseCase struct {
	repo domain.Repository
}

func NewCreateDeliveryUseCase(repo domain.Repository) *CreateDeliveryUseCase {
	return &CreateDeliveryUseCase{
		repo: repo,
	}
}

func (uc *CreateDeliveryUseCase) Execute(ctx context.Context, orderID, customerID, address string) (*domain.Delivery, error) {
	delivery := domain.NewDelivery(orderID, customerID, address)
	if err := uc.repo.Create(ctx, delivery); err != nil {
		return nil, err
	}
	return delivery, nil
} 