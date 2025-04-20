package repository

import (
	"merchant_inventory/config"
	"merchant_inventory/entities"
)

func Save(product entities.Product) (*entities.Product, error) {
	result := config.Database.Db.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func Update(product entities.Product) (*entities.Product, error) {
	result := config.Database.Db.Updates(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func FindAll(pageable Pageable) (Pageable, error) {
	products := []entities.Product{}

	result := Paginate(products, pageable, config.Database.Db)

	return result, nil
}

func FindByName(name string) ([]entities.Product, error) {
	products := []entities.Product{}

	result := config.Database.Db.Find(&products).
		Where("name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func FindBySku(sku string) (*entities.Product, error) {
	var product entities.Product

	result := config.Database.Db.First(&product, "sku = ?", sku)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func FindByUuid(uuid string) (*entities.Product, error) {
	var product entities.Product

	result := config.Database.Db.First(&product, "uuid = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
