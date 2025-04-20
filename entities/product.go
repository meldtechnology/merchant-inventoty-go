package entities

type Product struct {
	Id              int
	Uuid            string
	Name            string
	Description     string
	Sku             string
	Price           float64
	QuantityInStock int
	ReorderLevel    int
}
