package service

import (
	"log"
	purchaseOrderConverter "merchant_inventory/converter/purchase_order"
	supplierConverter "merchant_inventory/converter/supplier"
	"merchant_inventory/entities"
	"merchant_inventory/model"
	productRepository "merchant_inventory/repository"
	purchaseOrderRepository "merchant_inventory/repository/purchase_order"
	supplierRepository "merchant_inventory/repository/supplier"
	"merchant_inventory/util"
)

func AddPurchase(newPurchaseOrder model.PurchaseOrder) (*model.PurchaseOrder, error) {
	log.Println("Adding new Purchase Order...")

	//Get Supplier
	supplier, err := supplierRepository.FindByUuid(newPurchaseOrder.SupplierUuid)
	if err != nil {
		return nil, err
	}

	savedPurchaseOrder, err := purchaseOrderRepository.Save(purchaseOrderConverter.ToEntity(newPurchaseOrder, *supplier))
	if err != nil {
		return nil, err
	}

	// TODO: Implement async call to create a PURCHASE_ORDER Transaction in the inventory service
	supplierModel := supplierConverter.ToModel(*supplier)
	purchaseItems := createItems(newPurchaseOrder.Items, *savedPurchaseOrder)
	purchaseOrderModel := purchaseOrderConverter.ToModel(*savedPurchaseOrder, supplierModel, purchaseItems)
	return &purchaseOrderModel, nil
}

func UpdatePurchase(uuid string, status string) (*model.PurchaseOrder, error) {
	log.Println("Updating new Purchase Order...")

	purchaseOrderAndItems, err := purchaseOrderRepository.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	purchaseOrder := updatePurchaseOrderWithStatus(purchaseOrderAndItems.PurchaseOrder, status)

	_, err = purchaseOrderRepository.UpdatePurchaseOrder(purchaseOrder)
	// TODO: Implement async call here to update the product quantity in stock using goroutine when status is FULFILLED
	purchaseOrderAndItems.PurchaseOrder.PurchaseOrder = purchaseOrder
	savedPurchaseOrder := constructItem(*purchaseOrderAndItems)
	return &savedPurchaseOrder, nil
}

func GetPurchaseOrders(pageable productRepository.Pageable) (*productRepository.Pageable, error) {
	orderSupplierPaged, err := purchaseOrderRepository.FindAll(&pageable)

	if err != nil {
		return nil, err
	}

	purchaseOrders := constructItems(orderSupplierPaged)
	purchaseOrderPaged := productRepository.PostPaginate(purchaseOrders, &pageable)

	return &purchaseOrderPaged, nil
}

func GetPurchaseOrderByUuid(uuid string) (*model.PurchaseOrder, error) {
	purchaseOrderAndItems, err := purchaseOrderRepository.FindByUuid(uuid)

	if err != nil {
		return nil, err
	}

	purchaseOrders := constructItem(*purchaseOrderAndItems)

	return &purchaseOrders, nil
}

func GetStatues() ([]string, error) {
	var STATUSES []string
	var LAST_STATUS = int(util.FULFILLED)

	for index := 1; index <= LAST_STATUS; index = index + 1 {
		STATUSES = append(STATUSES, util.StatusName[util.STATUS_TYPE(index)])
	}

	return STATUSES, nil
}

func createItems(items []model.PurchaseOrderItem, purchaseOrder entities.PurchaseOrder) []model.PurchaseOrderItem {
	var purchaseOrderItems []entities.PurchaseOrderItem
	var mPurchaseOrderItems []model.PurchaseOrderItem
	cachedProduct := map[int]entities.Product{}

	for _, purchaseItems := range items {
		product := getProduct(purchaseItems.ProductId)
		cachedProduct[product.Id] = product
		purchaseOrderItems = append(purchaseOrderItems,
			purchaseOrderConverter.ToItemEntity(purchaseItems, purchaseOrder, product))
	}

	_, err := purchaseOrderRepository.SaveItems(purchaseOrderItems)
	if err != nil {
		return nil
	}

	// Convert to model
	for _, item := range purchaseOrderItems {
		product := cachedProduct[item.ProductId]
		mPurchaseOrderItems = append(mPurchaseOrderItems, purchaseOrderConverter.ToItemModel(item, product))
	}

	return mPurchaseOrderItems

}

func constructItems(purchaseAndItem []entities.PurchaseOrderAndItems) []model.PurchaseOrder {
	var purchaseOrders []model.PurchaseOrder

	for _, po := range purchaseAndItem {
		purchaseOrders = append(purchaseOrders, constructItem(po))
	}

	return purchaseOrders
}

func constructItem(purchaseAndItem entities.PurchaseOrderAndItems) model.PurchaseOrder {
	return purchaseOrderConverter.ToPurchaseAndItemModel(purchaseAndItem)
}

func getProduct(productUuid string) entities.Product {
	product, _ := productRepository.FindByUuid(productUuid)
	return *product
}

func updatePurchaseOrderWithStatus(order entities.PurchaseOrderSupplier, status string) entities.PurchaseOrder {
	order.PurchaseOrder.Status = status
	return order.PurchaseOrder
}
