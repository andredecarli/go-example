package customer

import "github.com/andredecarli/go-example/internal/domain/customer"

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (req CreateCustomerRequest) ToEntity() *customer.Customer {
	return &customer.Customer{
		Name:  req.Name,
		Email: req.Email,
	}
}
