package domain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, delivery *Delivery) error

	GetByID(ctx context.Context, id string) (*Delivery, error)

	Update(ctx context.Context, delivery *Delivery) error

	Delete(ctx context.Context, id string) error

	List(ctx context.Context, filter map[string]interface{}) ([]*Delivery, error)
} 