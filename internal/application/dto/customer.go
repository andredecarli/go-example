// Package dto provides data transfer objects for the application.
package dto

type CreateCustomerInput struct {
	Name  string
	Email string
}

type UpdateCustomerInput struct {
	ID    string
	Name  *string
	Email *string
}
