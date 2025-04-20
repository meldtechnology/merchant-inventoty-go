package model

type PurchaseOrderItem struct {
	Uuid        string  `json:"uuid"`
	ProductId   string  `json:"productId"`
	ProductName string  `json:"productName,omitempty"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}
