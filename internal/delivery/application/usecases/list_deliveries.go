package usecases

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type ListDeliveriesUseCase struct {
	repo domain.Repository
}

func NewListDeliveriesUseCase(repo domain.Repository) *ListDeliveriesUseCase {
	return &ListDeliveriesUseCase{
		repo: repo,
	}
}

func (uc *ListDeliveriesUseCase) Execute(ctx context.Context, filter map[string]interface{}) ([]*domain.Delivery, error) {
	return uc.repo.List(ctx, filter)
} 