package usecases

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type GetDeliveryUseCase struct {
	repo domain.Repository
}

func NewGetDeliveryUseCase(repo domain.Repository) *GetDeliveryUseCase {
	return &GetDeliveryUseCase{
		repo: repo,
	}
}

func (uc *GetDeliveryUseCase) Execute(ctx context.Context, id string) (*domain.Delivery, error) {
	return uc.repo.GetByID(ctx, id)
} 