package customer

import (
	"github.com/andredecarli/go-example/internal/domain/base"
	"github.com/andredecarli/go-example/internal/domain/customer"
)

func toEntity(model *customerModel) *customer.Customer {
	return &customer.Customer{
		Entity: base.Entity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:  model.Name,
		Email: model.Email,
	}
}

func toModel(entity *customer.Customer) *customerModel {
	return &customerModel{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
