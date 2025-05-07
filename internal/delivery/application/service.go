package application

import (
	"context"
	"logistic-system/internal/delivery/domain"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateDelivery(ctx context.Context, orderID, customerID, address string) (*domain.Delivery, error) {
	delivery := domain.NewDelivery(orderID, customerID, address)
	if err := s.repo.Create(ctx, delivery); err != nil {
		return nil, err
	}
	return delivery, nil
}

func (s *Service) GetDelivery(ctx context.Context, id string) (*domain.Delivery, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) UpdateDeliveryStatus(ctx context.Context, id string, status domain.Status) error {
	delivery, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := delivery.UpdateStatus(status); err != nil {
		return err
	}

	return s.repo.Update(ctx, delivery)
}

func (s *Service) DeleteDelivery(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) ListDeliveries(ctx context.Context, filter map[string]interface{}) ([]*domain.Delivery, error) {
	return s.repo.List(ctx, filter)
} 