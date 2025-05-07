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
	ID          string
	OrderID     string
	CustomerID  string
	Address     string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
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