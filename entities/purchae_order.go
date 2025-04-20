package entities

import "time"

type PurchaseOrder struct {
	Id           int
	Uuid         string
	SupplierId   int
	PurchaseDate time.Time
	Status       string
}
