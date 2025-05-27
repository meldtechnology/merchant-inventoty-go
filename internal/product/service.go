package product

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/meldtechnology/merchant-inventory-go/internal/entity"
	"github.com/meldtechnology/merchant-inventory-go/pkg/log"
)

// Service encapsulates use-cases logic for products.
type Service interface {
	Get(ctx context.Context, id string) (Product, error)
	Query(ctx context.Context, offset, limit int) ([]Product, error)
	Count(ctx context.Context) (int64, error)
	Create(ctx context.Context, input CreateProductRequest) (Product, error)
	Update(ctx context.Context, id string, input UpdateProductRequest) (Product, error)
	Delete(ctx context.Context, id string) (Product, error)
}

// Product represents the data about an product.
type Product struct {
	entity.Product
}

// CreateProductRequest represents an product creation request.
type CreateProductRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Sku             string  `json:"sku"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantityInStock"`
	ReorderLevel    int     `json:"reorderLevel"`
}

// Validate validates the CreateProductRequest fields.
func (m CreateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Price, validation.Required, validation.Min(10)),
	)
}

// UpdateProductRequest represents an product update request.
type UpdateProductRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantityInStock"`
	ReorderLevel    int     `json:"reorderLevel"`
}

// Validate validates the CreateProductRequest fields.
func (m UpdateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Price, validation.Required, validation.Min(10)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new product service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the product with the given product ID.
func (s service) Get(ctx context.Context, id string) (Product, error) {
	product, err := s.repo.Get(ctx, id)
	if err != nil {
		return Product{}, err
	}
	return Product{product}, nil
}

// Create creates a new product.
func (s service) Create(ctx context.Context, req CreateProductRequest) (Product, error) {
	if err := req.Validate(); err != nil {
		return Product{}, err
	}
	uuid := entity.GenerateID()
	err := s.repo.Create(ctx, entity.Product{
		Uuid:            uuid,
		Name:            req.Name,
		Sku:             req.Sku,
		Description:     req.Description,
		Price:           req.Price,
		QuantityInStock: req.QuantityInStock,
		ReorderLevel:    req.ReorderLevel,
	})
	if err != nil {
		return Product{}, err
	}
	return s.Get(ctx, uuid)
}

// Update updates the product with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateProductRequest) (Product, error) {
	if err := req.Validate(); err != nil {
		return Product{}, err
	}

	product, err := s.Get(ctx, id)
	if err != nil {
		return product, err
	}
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.QuantityInStock = req.QuantityInStock
	product.ReorderLevel = req.ReorderLevel

	if err := s.repo.Update(ctx, product.Product); err != nil {
		return product, err
	}
	return product, nil
}

// Delete deletes the product with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Product, error) {
	product, err := s.Get(ctx, id)
	if err != nil {
		return Product{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Product{}, err
	}
	return product, nil
}

// Count returns the number of product.
func (s service) Count(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

// Query returns the product with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Product, error) {
	if (offset - 1) < 0 {
		return nil, errors.New("invalid page value")
	}
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Product{}
	for _, item := range items {
		result = append(result, Product{item})
	}
	return result, nil
}
