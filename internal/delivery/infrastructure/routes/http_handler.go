package interfaces

import (
	"encoding/json"
	"logistic-system/internal/delivery/application"
	"logistic-system/internal/delivery/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/deliveries", h.CreateDelivery).Methods("POST")
	r.HandleFunc("/deliveries/{id}", h.GetDelivery).Methods("GET")
	r.HandleFunc("/deliveries/{id}/status", h.UpdateDeliveryStatus).Methods("PUT")
	r.HandleFunc("/deliveries/{id}", h.DeleteDelivery).Methods("DELETE")
	r.HandleFunc("/deliveries", h.ListDeliveries).Methods("GET")
}

type CreateDeliveryRequest struct {
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	Address    string `json:"address"`
}

func (h *Handler) CreateDelivery(w http.ResponseWriter, r *http.Request) {
	var req CreateDeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	delivery, err := h.service.CreateDelivery(r.Context(), req.OrderID, req.CustomerID, req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(delivery)
}

func (h *Handler) GetDelivery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	delivery, err := h.service.GetDelivery(r.Context(), id)
	if err != nil {
		if _, ok := err.(*domain.DomainError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(delivery)
}

type UpdateDeliveryStatusRequest struct {
	Status domain.Status `json:"status"`
}

func (h *Handler) UpdateDeliveryStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateDeliveryStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateDeliveryStatus(r.Context(), id, req.Status); err != nil {
		if _, ok := err.(*domain.DomainError); ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteDelivery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteDelivery(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListDeliveries(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]interface{})
	if status := r.URL.Query().Get("status"); status != "" {
		filter["status"] = domain.Status(status)
	}
	if customerID := r.URL.Query().Get("customer_id"); customerID != "" {
		filter["customer_id"] = customerID
	}

	deliveries, err := h.service.ListDeliveries(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deliveries)
} 