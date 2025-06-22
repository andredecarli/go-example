// Package services provides an implementation of the services.
package services

import (
	"errors"

	"github.com/andredecarli/go-example/internal/application/dto"
	"github.com/andredecarli/go-example/internal/domain/entities"
)

type CustomerRepository interface {
	Create(customer *entities.Customer) (*entities.Customer, error)
	FindAll() ([]*entities.Customer, error)
	FindByID(id string) (*entities.Customer, error)
	FindByEmail(email string) (*entities.Customer, error)
	Update(customer *entities.Customer) (*entities.Customer, error)
	Delete(id string) error
}

type customerService struct {
	customerRepository CustomerRepository
}

func NewCustomerService(customerRepository CustomerRepository) *customerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (s *customerService) Create(input dto.CreateCustomerInput) (*entities.Customer, error) {
	customer, err := s.customerRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if customer != nil {
		return nil, ErrCustomerAlreadyExists
	}
	customer = &entities.Customer{
		Name:  input.Name,
		Email: input.Email,
	}
	return s.customerRepository.Create(customer)
}

func (s *customerService) List() ([]*entities.Customer, error) {
	return s.customerRepository.FindAll()
}

func (s *customerService) FindByID(id string) (*entities.Customer, error) {
	customer, err := s.customerRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, ErrCustomerNotFound
	}
	return customer, nil
}

func (s *customerService) Update(input dto.UpdateCustomerInput) (*entities.Customer, error) {
	customer, err := s.customerRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, ErrCustomerNotFound
	}
	if input.Name != nil {
		customer.Name = *input.Name
	}
	if input.Email != nil {
		customer.Email = *input.Email
	}
	return s.customerRepository.Update(customer)
}

func (s *customerService) Delete(id string) error {
	return s.customerRepository.Delete(id)
}

var ErrCustomerAlreadyExists = errors.New("customer already exists")
var ErrCustomerNotFound = errors.New("customer not found")
