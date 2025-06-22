package customer

import (
	"github.com/andredecarli/go-example/internal/domain/base"
	"github.com/andredecarli/go-example/internal/domain/customer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toModel(entity *customer.Customer) *Customer {
	model := &Customer{
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.ID != "" {
		if objID, err := primitive.ObjectIDFromHex(entity.ID); err == nil {
			model.ID = objID
		}
	}

	return model
}

func toEntity(model *Customer) *customer.Customer {
	entity := &customer.Customer{
		Entity: base.Entity{
			ID:        model.ID.Hex(),
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:  model.Name,
		Email: model.Email,
	}
	return entity
}
