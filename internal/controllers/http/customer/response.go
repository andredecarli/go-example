package customer

import (
	"time"

	"github.com/andredecarli/go-example/internal/domain/customer"
)

type CustomerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (res *CustomerResponse) FromEntity(customer *customer.Customer) {
	res.ID = customer.ID
	res.Name = customer.Name
	res.Email = customer.Email
	res.CreatedAt = customer.CreatedAt
	res.UpdatedAt = customer.UpdatedAt
}
