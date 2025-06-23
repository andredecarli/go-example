// Package customer provides a controller for the customer domain.
package customer

import (
	"context"
	"errors"
	"log"

	"github.com/andredecarli/go-example/internal/domain/customer"
)

type CustomerService interface {
	Create(ctx context.Context, input *customer.Customer) (*customer.Customer, error)
}

type CustomerController struct {
	customerService CustomerService
}

func NewCustomerController(customerService CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) Create(ctx context.Context, userName string, userEmail string) {
	if userName == "" {
		log.Println("Error creating customer:", ErrInvalidName)
	}
	if userEmail == "" {
		log.Println("Error creating customer:", ErrInvalidEmail)
	}

	input := &customer.Customer{
		Name:  userName,
		Email: userEmail,
	}

	customer, err := c.customerService.Create(ctx, input)
	if err != nil {
		log.Println("Error creating customer:", err)
	}
	log.Println("Customer created:", customer)
}

var ErrInvalidName = errors.New("name is required")
var ErrInvalidEmail = errors.New("email is required")
