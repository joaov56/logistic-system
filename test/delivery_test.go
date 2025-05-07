package test

import (
	"logistic-system/internal/delivery/domain"
	"testing"
	"time"
)

func TestNewDelivery(t *testing.T) {
	orderID := "123"
	customerID := "456"
	address := "123 Main St"

	delivery := domain.NewDelivery(orderID, customerID, address)

	if delivery.OrderID != orderID {
		t.Errorf("Expected OrderID %s, got %s", orderID, delivery.OrderID)
	}

	if delivery.CustomerID != customerID {
		t.Errorf("Expected CustomerID %s, got %s", customerID, delivery.CustomerID)
	}

	if delivery.Address != address {
		t.Errorf("Expected Address %s, got %s", address, delivery.Address)
	}

	if delivery.Status != domain.StatusPending {
		t.Errorf("Expected Status %s, got %s", domain.StatusPending, delivery.Status)
	}

	if delivery.DeliveredAt != nil {
		t.Error("Expected DeliveredAt to be nil")
	}
}

func TestUpdateStatus(t *testing.T) {
	delivery := domain.NewDelivery("123", "456", "123 Main St")

	// Test valid status update
	err := delivery.UpdateStatus(domain.StatusInTransit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if delivery.Status != domain.StatusInTransit {
		t.Errorf("Expected Status %s, got %s", domain.StatusInTransit, delivery.Status)
	}

	// Test invalid status
	err = delivery.UpdateStatus("INVALID_STATUS")
	if err == nil {
		t.Error("Expected error for invalid status")
	}

	// Test delivered status sets DeliveredAt
	err = delivery.UpdateStatus(domain.StatusDelivered)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if delivery.DeliveredAt == nil {
		t.Error("Expected DeliveredAt to be set")
	}
}

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   domain.Status
		expected bool
	}{
		{"Pending", domain.StatusPending, true},
		{"InTransit", domain.StatusInTransit, true},
		{"Delivered", domain.StatusDelivered, true},
		{"Failed", domain.StatusFailed, true},
		{"Invalid", "INVALID", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidStatus(tt.status)
			if result != tt.expected {
				t.Errorf("Expected %v for status %s, got %v", tt.expected, tt.status, result)
			}
		})
	}
}

// Helper function to test status validation
func isValidStatus(s domain.Status) bool {
	switch s {
	case domain.StatusPending, domain.StatusInTransit, domain.StatusDelivered, domain.StatusFailed:
		return true
	default:
		return false
	}
} 