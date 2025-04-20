package _customer

import (
	"merchant_inventory/config"
	"merchant_inventory/entities"
	"merchant_inventory/repository"
)

func Save(customer entities.Customer) (*entities.Customer, error) {
	result := config.Database.Db.Create(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func Update(customer entities.Customer) (*entities.Customer, error) {
	result := config.Database.Db.Updates(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func FindAll(pageable repository.Pageable) (repository.Pageable, error) {
	customers := []entities.Customer{}

	result := repository.Paginate(customers, pageable, config.Database.Db)

	return result, nil
}

func FindByPhone(phone string) (*entities.Customer, error) {
	var customer entities.Customer

	result := config.Database.Db.First(&customer, "phone = ?", phone)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func FindByUuid(uuid string) (*entities.Customer, error) {
	var customer entities.Customer

	result := config.Database.Db.First(&customer, "uuid = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}
