package purchase_order

import (
	"github.com/google/uuid"
	"merchant_inventory/entities"
	"merchant_inventory/model"
	"time"
)

func ToEntity(model model.PurchaseOrder, supplier entities.Supplier) entities.PurchaseOrder {
	return entities.PurchaseOrder{
		Uuid:         checkForUuid(model.Uuid),
		SupplierId:   supplier.Id,
		PurchaseDate: useNow(),
		Status:       model.Status,
	}
}

func ToModel(entity entities.PurchaseOrder, supplier model.Supplier, poItems []model.PurchaseOrderItem) model.PurchaseOrder {
	return model.PurchaseOrder{
		Uuid:         entity.Uuid,
		SupplierUuid: supplier.Uuid,
		SupplierName: supplier.Name,
		PurchaseDate: entity.PurchaseDate,
		Status:       entity.Status,
		Items:        poItems,
	}
}

func ToItemEntity(model model.PurchaseOrderItem,
	purchaseOrder entities.PurchaseOrder,
	product entities.Product) entities.PurchaseOrderItem {
	return entities.PurchaseOrderItem{
		Uuid:            checkForUuid(model.Uuid),
		PurchaseOrderId: purchaseOrder.Id,
		ProductId:       product.Id,
		Quantity:        model.Quantity,
		UnitPrice:       model.UnitPrice,
	}
}

func ToItemModel(entity entities.PurchaseOrderItem, product entities.Product) model.PurchaseOrderItem {
	return model.PurchaseOrderItem{
		Uuid:        entity.Uuid,
		ProductId:   product.Uuid,
		ProductName: product.Name,
		Quantity:    entity.Quantity,
		UnitPrice:   entity.UnitPrice,
	}
}

func ToPurchaseAndItemModel(orderAndItems entities.PurchaseOrderAndItems) model.PurchaseOrder {
	var items []model.PurchaseOrderItem
	for _, item := range orderAndItems.Items {
		items = append(items, ToItemModel(item, entities.Product{}))
	}
	return toPurchaseOrderModel(
		orderAndItems.PurchaseOrder.PurchaseOrder,
		orderAndItems.PurchaseOrder.Supplier,
		items)
}

func toPurchaseOrderModel(entity entities.PurchaseOrder, supplier entities.Supplier, poItems []model.PurchaseOrderItem) model.PurchaseOrder {
	return model.PurchaseOrder{
		Uuid:         entity.Uuid,
		SupplierUuid: supplier.Uuid,
		SupplierName: supplier.Name,
		PurchaseDate: entity.PurchaseDate,
		Status:       entity.Status,
		Items:        poItems,
	}
}

func checkForUuid(value string) string {
	if value == "" {
		value = uuid.New().String()
	}
	return value
}

func useNow() time.Time {
	return time.Now()
}
