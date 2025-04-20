package _supplier

import (
	"merchant_inventory/config"
	"merchant_inventory/entities"
	"merchant_inventory/repository"
)

func Save(supplier entities.Supplier) (*entities.Supplier, error) {
	result := config.Database.Db.Create(&supplier)

	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}

func Update(supplier entities.Supplier) (*entities.Supplier, error) {
	result := config.Database.Db.Updates(&supplier)

	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}

func FindAll(pageable repository.Pageable) (repository.Pageable, error) {
	suppliers := []entities.Supplier{}

	result := repository.Paginate(suppliers, pageable, config.Database.Db)

	return result, nil
}

func FindByUuid(uuid string) (*entities.Supplier, error) {
	var supplier entities.Supplier

	result := config.Database.Db.First(&supplier, "uuid = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}
