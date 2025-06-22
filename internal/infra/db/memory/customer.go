// Package memory provides a memory-based implementation of the database interface.
package memory

import (
	"errors"
	"time"

	"github.com/andredecarli/go-example/internal/domain/entities"

	"github.com/google/uuid"
)

type customerRepository struct {
	data map[string]entities.Customer
}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{
		data: make(map[string]entities.Customer),
	}
}

func (r *customerRepository) Create(customer *entities.Customer) (*entities.Customer, error) {
	customer.ID = uuid.NewString()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	r.data[customer.ID] = *customer
	return customer, nil
}

func (r *customerRepository) FindByID(id string) (*entities.Customer, error) {
	customer, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	return &customer, nil
}

func (r *customerRepository) FindByEmail(email string) (*entities.Customer, error) {
	for _, customer := range r.data {
		if customer.Email == email {
			return &customer, nil
		}
	}
	return nil, nil
}

func (r *customerRepository) FindAll() ([]*entities.Customer, error) {
	customers := make([]*entities.Customer, 0, len(r.data))
	for _, customer := range r.data {
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *customerRepository) Update(customer *entities.Customer) (*entities.Customer, error) {
	_, ok := r.data[customer.ID]
	if !ok {
		return nil, ErrCustomerNotFound
	}
	customer.UpdatedAt = time.Now()

	r.data[customer.ID] = *customer
	return customer, nil
}

func (r *customerRepository) Delete(id string) error {
	_, ok := r.data[id]
	if !ok {
		return ErrCustomerNotFound
	}
	delete(r.data, id)
	return nil
}

var ErrCustomerNotFound = errors.New("customer not found")
