package customer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/andredecarli/go-example/internal/domain/customer"
)

type CustomerService interface {
	Create(ctx context.Context, input *customer.Customer) (*customer.Customer, error)
}

type handler struct {
	service CustomerService
}

func NewHandler(service CustomerService) *handler {
	return &handler{service: service}
}

func (h *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input := req.ToEntity()

	output, err := h.service.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response CustomerResponse
	response.FromEntity(output)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
