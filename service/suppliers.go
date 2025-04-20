package service

import (
	"encoding/json"
	"log"
	"merchant_inventory/converter/supplier"
	"merchant_inventory/entities"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	_supplier "merchant_inventory/repository/supplier"
)

func AddSupplier(newSupplier model.Supplier) (*model.Supplier, error) {
	log.Println("Adding new Supplier ...")

	savedSupplier, err := _supplier.Save(supplier.ToEntity(newSupplier))
	if err != nil {
		return nil, err
	}

	supplierModel := supplier.ToModel(*savedSupplier)
	return &supplierModel, nil
}

func GetSuppliers(pageable repository.Pageable) (*repository.Pageable, error) {
	log.Println("Retrieving paginated Suppliers")

	suppliers, err := _supplier.FindAll(pageable)

	err = mapToSupplierModel(&suppliers)
	if err != nil {
		return nil, err
	}
	return &suppliers, nil
}

func GetSuppliersByUuid(uuid string) (*model.Supplier, error) {
	log.Println("Retrieving Supplier by uuid ", uuid)

	foundSupplier, err := _supplier.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	supplierModels := supplier.ToModel(*foundSupplier)
	return &supplierModels, nil
}

func UpdateSuppliersByUuid(uuid string, updatedSupplier model.Supplier) (*model.Supplier, error) {
	log.Println("Updating Supplier by uuid")

	// Find the Supplier given the sku
	foundSupplier, err := _supplier.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	savedSupplier, err := _supplier.Update(modifySupplier(updatedSupplier, *foundSupplier))
	if err != nil {
		return nil, err
	}

	supplierModel := supplier.ToModel(*savedSupplier)
	return &supplierModel, nil
}

func modifySupplier(model model.Supplier, entity entities.Supplier) entities.Supplier {
	entity.Name = model.Name
	entity.ContactInfo = model.ContactInfo
	return entity
}

func mapToSupplierModel(supplierPaged *repository.Pageable) error {
	marshal, err := json.Marshal(supplierPaged.Rows)
	domain := make([]entities.Supplier, supplierPaged.Limit)
	err = json.Unmarshal(marshal, &domain)
	supplierModels := supplier.ToModels(domain)
	supplierPaged.Rows = supplierModels
	return err
}
