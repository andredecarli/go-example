// Package customer provides a MongoDB implementation of the customer repository.
package customer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName = "customers"

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
