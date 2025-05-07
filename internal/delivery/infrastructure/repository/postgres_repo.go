package infrastructure

import (
	"context"
	"logistic-system/internal/delivery/domain"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, delivery *domain.Delivery) error {
	return r.db.WithContext(ctx).Create(delivery).Error
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*domain.Delivery, error) {
	var delivery domain.Delivery
	err := r.db.WithContext(ctx).First(&delivery, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.NewDomainError("delivery not found")
		}
		return nil, err
	}
	return &delivery, nil
}

func (r *PostgresRepository) Update(ctx context.Context, delivery *domain.Delivery) error {
	return r.db.WithContext(ctx).Save(delivery).Error
}

func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&domain.Delivery{}, "id = ?", id).Error
}

func (r *PostgresRepository) List(ctx context.Context, filter map[string]interface{}) ([]*domain.Delivery, error) {
	var deliveries []*domain.Delivery
	query := r.db.WithContext(ctx).Model(&domain.Delivery{})

	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Find(&deliveries).Error
	if err != nil {
		return nil, err
	}

	return deliveries, nil
} 