package model

import (
	"time"
)

type PurchaseOrder struct {
	Uuid         string              `json:"uuid"`
	SupplierUuid string              `json:"supplierId"`
	SupplierName string              `json:"supplierName"`
	PurchaseDate time.Time           `json:"purchaseDate"`
	Status       string              `json:"status"`
	Items        []PurchaseOrderItem `json:"items"`
}
