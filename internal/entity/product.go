package entity

type Product struct {
	Id              int     `json:"id"`
	Uuid            string  `json:"uuid"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Sku             string  `json:"sku"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantity_in_stock"`
	ReorderLevel    int     `json:"reorder_level"`
}
