package adding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddProducts(t *testing.T) {
	p := Product{
		Name:            "Test Product 1",
		Description:     "Doll One",
		Sku:             "OIUYTRE",
		QuantityInStock: 20,
		Price:           15000,
		ReorderLevel:    10,
	}

	mR := new(mockStorage)

	s := NewService(mR)

	err := s.AddProduct(p)
	assert.NoError(t, err)

}

type mockStorage struct {
	p Product
}

func (m *mockStorage) AddProduct(b Product) error {
	m.p = b

	return nil
}
