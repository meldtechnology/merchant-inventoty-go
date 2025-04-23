package listing

import "errors"

// ErrNotFound is used when a product could not be found.
var ErrNotFound = errors.New("product not found")

// Repository provides access to the product storage.
type Repository interface {
	// GetProduct returns the product with given ID.
	GetProduct(int) (Product, error)
	// GetProductSku returns the product with given Sku.
	GetProductSku(string) (Product, error)
	// GetAllProducts returns all products saved in storage.
	GetAllProducts(Pageable) (Pageable, error)
}

// Service provides product and review listing operations.
type Service interface {
	GetProducts(Pageable) (Pageable, error)
	GetProduct(int) (Product, error)
	GetProductSku(string) (Product, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetProducts returns all products
func (s *service) GetProducts(p Pageable) (Pageable, error) {
	return s.r.GetAllProducts(p)
}

// GetProduct returns a product
func (s *service) GetProduct(id int) (Product, error) {
	return s.r.GetProduct(id)
}

// GetProductSku returns product by the sku
func (s *service) GetProductSku(sku string) (Product, error) {
	return s.r.GetProductSku(sku)
}
