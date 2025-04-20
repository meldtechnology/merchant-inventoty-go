package converter

import (
	"github.com/google/uuid"
	"merchant_inventory/entities"
	"merchant_inventory/model"
)

func ToEntity(model model.Product) entities.Product {
	return entities.Product{
		Uuid:            checkForUuid(model.Uuid),
		Name:            model.Name,
		Description:     model.Description,
		Sku:             model.Sku,
		Price:           model.Price,
		QuantityInStock: model.QuantityInStock,
		ReorderLevel:    model.ReorderLevel,
	}
}

func ToModel(entity entities.Product) model.Product {
	return model.Product{
		Uuid:            entity.Uuid,
		Name:            entity.Name,
		Description:     entity.Description,
		Sku:             entity.Sku,
		Price:           entity.Price,
		QuantityInStock: entity.QuantityInStock,
		ReorderLevel:    entity.ReorderLevel,
	}
}

func ToModels(entities []entities.Product) []model.Product {
	var models []model.Product

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
