package adding

// Service provides product, customer, supplier adding operations.
type Service interface {
	AddProduct(Product) error
}

// Repository provides access to product, customer, supplier repository.
type Repository interface {
	// AddProduct saves a given product to the repository.
	AddProduct(Product) error
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddProduct persists the given product(s) to storage
func (s *service) AddProduct(p Product) error {
	// validate add for any duplicates
	_ = s.r.AddProduct(p)

	return nil
}
