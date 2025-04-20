package customer

import (
	"github.com/google/uuid"
	"merchant_inventory/entities"
	"merchant_inventory/model"
)

func ToEntity(model model.Customer) entities.Customer {
	return entities.Customer{
		Uuid:            checkForUuid(model.Uuid),
		Name:            model.Name,
		Email:           model.Email,
		Phone:           model.Phone,
		ShippingAddress: model.ShippingAddress,
	}
}

func ToModel(entity entities.Customer) model.Customer {
	return model.Customer{
		Uuid:            entity.Uuid,
		Name:            entity.Name,
		Email:           entity.Email,
		Phone:           entity.Phone,
		ShippingAddress: entity.ShippingAddress,
	}
}

func ToModels(entities []entities.Customer) []model.Customer {
	var models []model.Customer

	for _, entity := range entities {
		models = append(models, ToModel(entity))
	}

	return models
}

func checkForUuid(value string) string {
	if value == "" {
		value = uuid.New().String()
	}
	return value
}
