package usecases

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type UpdateDeliveryStatusUseCase struct {
	repo domain.Repository
}

func NewUpdateDeliveryStatusUseCase(repo domain.Repository) *UpdateDeliveryStatusUseCase {
	return &UpdateDeliveryStatusUseCase{
		repo: repo,
	}
}

func (uc *UpdateDeliveryStatusUseCase) Execute(ctx context.Context, id string, status domain.Status) error {
	delivery, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := delivery.UpdateStatus(status); err != nil {
		return err
	}

	return uc.repo.Update(ctx, delivery)
} 