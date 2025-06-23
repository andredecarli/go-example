// Package customer contains the memory implementation of the customer repository.
package customer

import "time"

type customerModel struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
