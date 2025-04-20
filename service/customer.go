package service

import (
	"encoding/json"
	"log"
	"merchant_inventory/converter/customer"
	"merchant_inventory/entities"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	"merchant_inventory/repository/customer"
)

func AddCustomer(newCustomer model.Customer) (*model.Customer, error) {
	log.Println("Adding new Customer...")

	savedCustomer, err := _customer.Save(customer.ToEntity(newCustomer))
	if err != nil {
		return nil, err
	}

	customerModel := customer.ToModel(*savedCustomer)
	return &customerModel, nil
}

func GetCustomers(pageable repository.Pageable) (*repository.Pageable, error) {
	log.Println("Retrieving paginated Customers")

	customers, err := _customer.FindAll(pageable)

	err = mapToCustomerModel(&customers)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func GetCustomersByPhone(phone string) (*model.Customer, error) {
	log.Println("Retrieving Customer by phone ", phone)

	foundCustomer, err := _customer.FindByPhone(phone)
	if err != nil {
		return nil, err
	}

	customerModels := customer.ToModel(*foundCustomer)
	return &customerModels, nil
}

func GetCustomersByUuid(uuid string) (*model.Customer, error) {
	log.Println("Retrieving Customer by uuid ", uuid)

	foundCustomer, err := _customer.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	customerModels := customer.ToModel(*foundCustomer)
	return &customerModels, nil
}

func UpdateCustomersByUuid(uuid string, updatedCustomer model.Customer) (*model.Customer, error) {
	log.Println("Updating Customer by uuid")

	// Find the Customer given the sku
	foundCustomer, err := _customer.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	savedCustomer, err := _customer.Update(modifyCustomer(updatedCustomer, *foundCustomer))
	if err != nil {
		return nil, err
	}

	customerModel := customer.ToModel(*savedCustomer)
	return &customerModel, nil
}

func modifyCustomer(model model.Customer, entity entities.Customer) entities.Customer {
	entity.Name = model.Name
	entity.Phone = model.Phone
	entity.Email = model.Email
	entity.ShippingAddress = model.ShippingAddress
	return entity
}

func mapToCustomerModel(customerPaged *repository.Pageable) error {
	marshal, err := json.Marshal(customerPaged.Rows)
	domain := make([]entities.Customer, customerPaged.Limit)
	err = json.Unmarshal(marshal, &domain)
	customerModels := customer.ToModels(domain)
	customerPaged.Rows = customerModels
	return err
}
