package entities

type PurchaseOrderItem struct {
	Id              int
	Uuid            string
	PurchaseOrderId int
	ProductId       int
	Quantity        int
	UnitPrice       float64
}
