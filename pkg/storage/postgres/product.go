package postgres

type Product struct {
	Id              int     `json:"id,omitempty"`
	Uuid            string  `json:"uuid,omitempty"`
	Name            string  `json:"name,omitempty"`
	Description     string  `json:"description,omitempty"`
	Sku             string  `json:"sku,omitempty"`
	Price           float64 `json:"price,omitempty"`
	QuantityInStock int     `json:"quantityInStock,omitempty"`
	ReorderLevel    int     `json:"reorderLevel,omitempty"`
}
