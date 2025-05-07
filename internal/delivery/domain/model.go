package domain

import (
	"time"
)

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusInTransit Status = "IN_TRANSIT"
	StatusDelivered Status = "DELIVERED"
	StatusFailed    Status = "FAILED"
)

type Delivery struct {
	ID          string    `gorm:"primaryKey;type:varchar(36)"`
	OrderID     string    `gorm:"type:varchar(36);not null"`
	CustomerID  string    `gorm:"type:varchar(36);not null"`
	Address     string    `gorm:"type:text;not null"`
	Status      Status    `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeliveredAt *time.Time
}

func NewDelivery(orderID, customerID, address string) *Delivery {
	now := time.Now()
	return &Delivery{
		OrderID:    orderID,
		CustomerID: customerID,
		Address:    address,
		Status:     StatusPending,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (d *Delivery) UpdateStatus(status Status) error {
	if !isValidStatus(status) {
		return ErrInvalidStatus
	}

	d.Status = status
	d.UpdatedAt = time.Now()

	if status == StatusDelivered {
		now := time.Now()
		d.DeliveredAt = &now
	}

	return nil
}

func isValidStatus(s Status) bool {
	switch s {
	case StatusPending, StatusInTransit, StatusDelivered, StatusFailed:
		return true
	default:
		return false
	}
}

var (
	ErrInvalidStatus = NewDomainError("invalid delivery status")
)

type DomainError struct {
	message string
}

func (e *DomainError) Error() string {
	return e.message
}

func NewDomainError(message string) error {
	return &DomainError{message: message}
} 