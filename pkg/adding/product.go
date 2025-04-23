package adding

type Product struct {
	Id              int     `json:"id"`
	Uuid            string  `json:"uuid"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Sku             string  `json:"sku"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantityInStock"`
	ReorderLevel    int     `json:"reorderLevel"`
}
