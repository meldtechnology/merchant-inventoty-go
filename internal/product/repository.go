package product

import (
	"context"
	"github.com/mrchantinevntory/internal/entity"
	"github.com/mrchantinevntory/pkg/dbcontext/progress"
	"github.com/mrchantinevntory/pkg/log"
)

// Repository encapsulates the logic to access products from the data source.
type Repository interface {
	// Get returns the product with the specified product ID.
	Get(ctx context.Context, id string) (entity.Product, error)
	// Count returns the number of products.
	Count(ctx context.Context) (int64, error)
	// Query returns the list of products with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Product, error)
	// Create saves a new product in the storage.
	Create(ctx context.Context, product entity.Product) error
	// Update updates the product with given ID in the storage.
	Update(ctx context.Context, product entity.Product) error
	// Delete removes the product with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists products in database
type repository struct {
	db     *progress.PgDb
	logger log.Logger
}

// NewRepository creates a new product repository
func NewRepository(db *progress.PgDb, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the product with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product
	err := r.db.With(ctx).Db.First(&product, id)
	return product, err.Error
}

// Create saves a new product record in the database.
// It returns the ID of the newly inserted product record.
func (r repository) Create(ctx context.Context, product entity.Product) error {
	return r.db.With(ctx).Db.Create(&product).Error
}

// Update saves the changes to an product in the database.
func (r repository) Update(ctx context.Context, product entity.Product) error {
	return r.db.With(ctx).Db.Save(&product).Error
}

// Delete deletes an product with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	product, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Db.Delete(&product).Error
}

// Count returns the number of the product records in the database.
func (r repository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.With(ctx).Db.Model(&entity.Product{}).Count(&count).Error
	return count, err
}

// Query retrieves the product records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.With(ctx).Db.
		Order("id").
		Offset(offset - 1).
		Limit(limit).
		Find(&products).
		Error

	return products, err
}
