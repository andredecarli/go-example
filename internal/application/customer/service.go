// Package customer contains the basic customer CRUD service
package customer

import (
	"context"
	"errors"

	"github.com/andredecarli/go-example/internal/domain/customer"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=customer

// CustomerRepository is the interface that wraps the Create method.
type CustomerRepository interface {
	Create(ctx context.Context, customer *customer.Customer) (*customer.Customer, error)
}

type service struct {
	repository CustomerRepository
}

func NewService(repository CustomerRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input *customer.Customer) (*customer.Customer, error) {
	if input.Name == "" {
		return nil, ErrNameIsRequired
	}
	if input.Email == "" {
		return nil, ErrEmailIsRequired
	}

	return s.repository.Create(ctx, input)
}

var ErrNameIsRequired = errors.New("name is required")
var ErrEmailIsRequired = errors.New("email is required")
