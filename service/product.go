package service

import (
	"encoding/json"
	"log"
	"merchant_inventory/converter"
	"merchant_inventory/entities"
	"merchant_inventory/model"
	"merchant_inventory/repository"
)

func AddProduct(newProduct model.Product) (*model.Product, error) {
	log.Println("Adding new Product to Master Catalog")

	savedProduct, err := repository.Save(converter.ToEntity(newProduct))
	if err != nil {
		return nil, err
	}

	productModel := converter.ToModel(*savedProduct)
	return &productModel, nil
}

func GetProducts(pageable repository.Pageable) (*repository.Pageable, error) {
	log.Println("Retrieving paginated Products from Master Catalog")

	products, err := repository.FindAll(pageable)

	err = mapToModel(&products)
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func GetProductsBySku(sku string) (*model.Product, error) {
	log.Println("Retrieving Product by sku from Master Catalog ", sku)

	product, err := repository.FindBySku(sku)
	if err != nil {
		return nil, err
	}

	productModels := converter.ToModel(*product)
	return &productModels, nil
}

func GetProductsByUuid(uuid string) (*model.Product, error) {
	log.Println("Retrieving Product by sku from Master Catalog ", uuid)

	product, err := repository.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	productModels := converter.ToModel(*product)
	return &productModels, nil
}

func UpdateProductsByUuid(uuid string, updatedProduct model.Product) (*model.Product, error) {
	log.Println("Updating Product by sku in the Master Catalog")

	// Find the product given the sku
	foundProduct, err := repository.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	savedProduct, err := repository.Update(modifyProduct(updatedProduct, *foundProduct))
	if err != nil {
		return nil, err
	}

	productModel := converter.ToModel(*savedProduct)
	return &productModel, nil
}

func modifyProduct(model model.Product, entity entities.Product) entities.Product {
	entity.Name = model.Name
	entity.Price = model.Price
	entity.QuantityInStock = model.QuantityInStock
	entity.ReorderLevel = model.ReorderLevel
	return entity
}

func mapToModel(productPaged *repository.Pageable) error {
	marshal, err := json.Marshal(productPaged.Rows)
	domain := make([]entities.Product, productPaged.Limit)
	err = json.Unmarshal(marshal, &domain)
	productModels := converter.ToModels(domain)
	productPaged.Rows = productModels
	return err
}
