package _purchase_order

import (
	"errors"
	"merchant_inventory/config"
	"merchant_inventory/entities"
	"merchant_inventory/repository"
)

var SELECT_PURCHASE_ORDER_SUPPLIER = "SELECT p.*, s.* from purchase_orders p LEFT JOIN suppliers s ON p.supplier_id = s.id order by p.id desc OFFSET ? LIMIT ?"
var SELECT_PURCHASE_ORDER_SUPPLIER_BY_UUID = "SELECT p.*, s.* from purchase_orders p LEFT JOIN suppliers s ON p.supplier_id = s.id WHERE p.uuid = ?"

func Save(purchaseOrder entities.PurchaseOrder) (*entities.PurchaseOrder, error) {
	result := config.Database.Db.Create(&purchaseOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return &purchaseOrder, nil
}

func SaveItems(purchaseOrder []entities.PurchaseOrderItem) ([]entities.PurchaseOrderItem, error) {
	result := config.Database.Db.Create(&purchaseOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return purchaseOrder, nil
}

func UpdatePurchaseOrder(purchaseOrder entities.PurchaseOrder) (*entities.PurchaseOrder, error) {
	result := config.Database.Db.Updates(&purchaseOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return &purchaseOrder, nil
}

func FindAll(pageable *repository.Pageable) ([]entities.PurchaseOrderAndItems, error) {
	var purchaseAndItems []entities.PurchaseOrderAndItems
	purchaseOrderSupplier := findPurchaseOrderSupplier(pageable)

	for _, orderSupplier := range purchaseOrderSupplier {
		purchaseAndItems = append(purchaseAndItems, findWithItemsInclusive(orderSupplier))
	}

	return purchaseAndItems, nil
}

func findPurchaseOrderSupplier(pageable *repository.Pageable) []entities.PurchaseOrderSupplier {
	var purchaseOrderSupplier []entities.PurchaseOrderSupplier

	result := repository.PrePaginate([]entities.PurchaseOrder{}, pageable, config.Database.Db).
		Raw(SELECT_PURCHASE_ORDER_SUPPLIER, pageable.Page, pageable.Limit).
		Scan(&purchaseOrderSupplier)

	if result.Error != nil {
		return nil
	}

	return purchaseOrderSupplier
}

func findWithItemsInclusive(orderSupplier entities.PurchaseOrderSupplier) entities.PurchaseOrderAndItems {
	var items []entities.PurchaseOrderItem
	config.Database.Db.Find(&items).
		Where("purchase_order_id = ?", orderSupplier.PurchaseOrder.Id)

	return entities.PurchaseOrderAndItems{orderSupplier, items}
}

func FindByUuid(uuid string) (*entities.PurchaseOrderAndItems, error) {
	var purchaseOrderSupplier entities.PurchaseOrderSupplier

	result := config.Database.Db.Raw(SELECT_PURCHASE_ORDER_SUPPLIER_BY_UUID, uuid).
		Scan(&purchaseOrderSupplier)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("No Record Found")
	}

	purchaseOrder := findWithItemsInclusive(purchaseOrderSupplier)

	return &purchaseOrder, nil
}
