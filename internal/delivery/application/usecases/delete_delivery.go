package usecases

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type DeleteDeliveryUseCase struct {
	repo domain.Repository
}

func NewDeleteDeliveryUseCase(repo domain.Repository) *DeleteDeliveryUseCase {
	return &DeleteDeliveryUseCase{
		repo: repo,
	}
}

func (uc *DeleteDeliveryUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
} 