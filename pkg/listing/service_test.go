package listing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProducts(t *testing.T) {
	p1 := Product{
		Id:              1,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "PLO0987UJI",
		QuantityInStock: 15,
		Price:           35000,
		ReorderLevel:    5,
	}

	p2 := Product{
		Id:              2,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "PLO0987UJI",
		QuantityInStock: 15,
		Price:           35000,
		ReorderLevel:    5,
	}

	mR := new(mockStorage)

	mR.p = append(mR.p, p1, p2)

	ps, _ := mR.GetAllProducts()
	assert.NotNil(t, ps)
	assert.Len(t, ps.Rows, 2)
}

func TestGetProduct(t *testing.T) {
	p1 := Product{
		Id:              1,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "PLO0987UJI",
		QuantityInStock: 15,
		Price:           35000,
		ReorderLevel:    5,
	}

	p2 := Product{
		Id:              2,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "PLO0987UJI",
		QuantityInStock: 15,
		Price:           35000,
		ReorderLevel:    5,
	}

	mR := new(mockStorage)

	mR.p = append(mR.p, p1, p2)

	ps, _ := mR.GetProduct(2)
	assert.NotNil(t, ps)
}

func TestGetProductSku(t *testing.T) {
	p1 := Product{
		Id:              1,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "PLO0987UJI",
		QuantityInStock: 15,
		Price:           35000,
		ReorderLevel:    5,
	}

	p2 := Product{
		Id:              2,
		Name:            "Product Test 1",
		Description:     "Diary Product Test 1",
		Sku:             "KLO0987UJI",
		QuantityInStock: 55,
		Price:           65000,
		ReorderLevel:    15,
	}

	mR := new(mockStorage)

	mR.p = append(mR.p, p1, p2)

	ps, _ := mR.GetProductSku("PLO0987UJI")
	assert.NotNil(t, ps)
}

type mockStorage struct {
	p []Product
}

func (m *mockStorage) GetProduct(id int) (Product, error) {
	var prd Product

	for _, pr := range m.p {
		if pr.Id == id {
			prd = pr
		}
	}

	return prd, nil
}

func (m *mockStorage) GetProductSku(s string) (Product, error) {
	var prd Product

	for _, pr := range m.p {
		if pr.Sku == s {
			prd = pr
		}
	}

	return prd, nil
}

func (m *mockStorage) GetAllProducts() (Pageable, error) {
	prds := []Product{}

	for _, prd := range m.p {
		pr := Product{
			Id:              prd.Id,
			Name:            prd.Name,
			Description:     prd.Description,
			Sku:             prd.Sku,
			QuantityInStock: prd.QuantityInStock,
			Price:           prd.Price,
			ReorderLevel:    prd.ReorderLevel,
		}
		prds = append(prds, pr)
	}

	total := int64(len(prds))
	pg := Pageable{Page: 1, Limit: 10, TotalRows: total, Rows: prds}

	return pg, nil
}
