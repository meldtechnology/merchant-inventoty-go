package supplier

import (
	"github.com/google/uuid"
	"merchant_inventory/entities"
	"merchant_inventory/model"
)

func ToEntity(model model.Supplier) entities.Supplier {
	return entities.Supplier{
		Uuid:        checkForUuid(model.Uuid),
		Name:        model.Name,
		ContactInfo: model.ContactInfo,
	}
}

func ToModel(entity entities.Supplier) model.Supplier {
	return model.Supplier{
		Uuid:        entity.Uuid,
		Name:        entity.Name,
		ContactInfo: entity.ContactInfo,
	}
}

func ToModels(entities []entities.Supplier) []model.Supplier {
	var models []model.Supplier

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
