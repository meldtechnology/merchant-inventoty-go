package model

type Customer struct {
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	ShippingAddress string `json:"shippingAddress"`
}
