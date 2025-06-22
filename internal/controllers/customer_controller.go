// Package controllers
package controllers

import (
	"errors"

	"github.com/andredecarli/go-example/internal/application/dto"
	"github.com/andredecarli/go-example/internal/domain/entities"
)

type CustomerService interface {
	Create(input dto.CreateCustomerInput) (*entities.Customer, error)
	List() ([]*entities.Customer, error)
	FindByID(id string) (*entities.Customer, error)
	Update(input dto.UpdateCustomerInput) (*entities.Customer, error)
	Delete(id string) error
}

type CustomerController struct {
	customerService CustomerService
}

func NewCustomerController(customerService CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) Create(userName string, userEmail string) (*entities.Customer, error) {
	if userName == "" {
		return nil, ErrInvalidName
	}
	if userEmail == "" {
		return nil, ErrInvalidEmail
	}

	input := dto.CreateCustomerInput{
		Name:  userName,
		Email: userEmail,
	}

	customer, err := c.customerService.Create(input)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerController) List() ([]*entities.Customer, error) {
	return c.customerService.List()
}

func (c *CustomerController) FindByID(id string) (*entities.Customer, error) {
	return c.customerService.FindByID(id)
}

func (c *CustomerController) Update(userID string, userName string, userEmail string) (*entities.Customer, error) {
	input := dto.UpdateCustomerInput{
		ID: userID,
	}
	if userName != "" {
		input.Name = &userName
	}
	if userEmail != "" {
		input.Email = &userEmail
	}

	return c.customerService.Update(input)
}

func (c *CustomerController) Delete(id string) error {
	return c.customerService.Delete(id)
}

var ErrInvalidName = errors.New("name is required")
var ErrInvalidEmail = errors.New("email is required")
var ErrInvalidInput = errors.New("input is required")
