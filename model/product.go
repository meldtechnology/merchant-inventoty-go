package model

type Product struct {
	Uuid            string  `json:"uuid"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Sku             string  `json:"sku"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantityInHand"`
	ReorderLevel    int     `json:"reorderLevel"`
}
