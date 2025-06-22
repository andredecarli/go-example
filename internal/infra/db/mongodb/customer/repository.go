package customer

import (
	"context"

	"github.com/andredecarli/go-example/internal/domain/customer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection(collectionName)}
}

func (r *Repository) Create(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	model := toModel(customer)

	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	model.ID = result.InsertedID.(primitive.ObjectID)

	return toEntity(model), nil
}
