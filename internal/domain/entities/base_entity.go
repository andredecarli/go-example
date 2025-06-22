// Package entities
package entities

import "time"

type Entity struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}
