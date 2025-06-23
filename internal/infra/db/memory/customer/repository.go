package customer

import (
	"context"
	"time"

	"github.com/andredecarli/go-example/internal/domain/customer"
	"github.com/google/uuid"
)

type repository struct {
	customers map[string]*customerModel
}

func NewRepository() *repository {
	return &repository{customers: make(map[string]*customerModel)}
}

func (r *repository) Create(_ context.Context, customer *customer.Customer) (*customer.Customer, error) {
	model := toModel(customer)
	model.ID = uuid.NewString()
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	r.customers[model.ID] = model
	return toEntity(model), nil
}
