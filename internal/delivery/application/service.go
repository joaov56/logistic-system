package application

import (
	"context"
	"logistic-system/internal/delivery/application/usecases"
	"logistic-system/internal/delivery/domain"
)

type Service struct {
	createDeliveryUseCase      *usecases.CreateDeliveryUseCase
	getDeliveryUseCase         *usecases.GetDeliveryUseCase
	updateDeliveryStatusUseCase *usecases.UpdateDeliveryStatusUseCase
	deleteDeliveryUseCase      *usecases.DeleteDeliveryUseCase
	listDeliveriesUseCase      *usecases.ListDeliveriesUseCase
}

func NewService(repo domain.Repository) *Service {
	return &Service{
		createDeliveryUseCase:      usecases.NewCreateDeliveryUseCase(repo),
		getDeliveryUseCase:         usecases.NewGetDeliveryUseCase(repo),
		updateDeliveryStatusUseCase: usecases.NewUpdateDeliveryStatusUseCase(repo),
		deleteDeliveryUseCase:      usecases.NewDeleteDeliveryUseCase(repo),
		listDeliveriesUseCase:      usecases.NewListDeliveriesUseCase(repo),
	}
}

func (s *Service) CreateDelivery(ctx context.Context, orderID, customerID, address string) (*domain.Delivery, error) {
	return s.createDeliveryUseCase.Execute(ctx, orderID, customerID, address)
}

func (s *Service) GetDelivery(ctx context.Context, id string) (*domain.Delivery, error) {
	return s.getDeliveryUseCase.Execute(ctx, id)
}

func (s *Service) UpdateDeliveryStatus(ctx context.Context, id string, status domain.Status) error {
	return s.updateDeliveryStatusUseCase.Execute(ctx, id, status)
}

func (s *Service) DeleteDelivery(ctx context.Context, id string) error {
	return s.deleteDeliveryUseCase.Execute(ctx, id)
}

func (s *Service) ListDeliveries(ctx context.Context, filter map[string]interface{}) ([]*domain.Delivery, error) {
	return s.listDeliveriesUseCase.Execute(ctx, filter)
} 