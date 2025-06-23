// Package base contains the base entity for all entities
package base

import "time"

type Entity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
