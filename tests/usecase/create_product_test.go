package usecase

import (
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	product := &productEntities.Product{
		Title:       "product title",
		Description: "product description",
		Price:       100,
		Rating:      5,
		CategoryID:  1,
		VendorID:    1,
		LocationID:  1,
		BrandID:     1,
		Stock:       10,
	}
	repo := new(mock_product.MockProductRepository)
	productUseCase := productUseCase.NewProductUseCase(repo)

	err := productUseCase.CreateProduct(product)
	if err != nil {
		t.Errorf("Error while creating product, got: %v", err)
	}
	if product.ID == 0 {
		t.Errorf("Product ID is not set, got: %d", product.ID)
	}
}
