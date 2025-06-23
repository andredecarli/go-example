// Package customer contains the customer entity
package customer

import (
	"fmt"

	"github.com/andredecarli/go-example/internal/domain/base"
	"github.com/andredecarli/go-example/pkg/util"
)

type Customer struct {
	base.Entity

	Name  string
	Email string
}

func (c Customer) String() string {
	return fmt.Sprintf(
		"Customer{ID: %s, Name: %s, Email: %s, CreatedAt: %s, UpdatedAt: %s}",
		c.ID[:8],
		c.Name,
		c.Email,
		util.FormatTime(c.CreatedAt),
		util.FormatTime(c.UpdatedAt),
	)
}
