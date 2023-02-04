package usecase

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	mockProductRepository := new(mock_product.MockProductRepository)
	productUseCase := productUseCase.NewProductUseCase(mockProductRepository)

	mockProduct := &productEntities.Product{
		ID:          1,
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

	result, err := productUseCase.GetByID(context.Background(), 1)
	if err != nil {
		t.Errorf("Error while getting product, got: %v", err)
	}

	if result == nil {
		t.Errorf("No product found, got: %v", result)
	}

	if result.ID != mockProduct.ID {
		t.Errorf("Product ID does not match, expected: %d, got: %d", mockProduct.ID, result.ID)
	}

	if result.Title != mockProduct.Title {
		t.Errorf("Product title does not match, expected: %s, got: %s", mockProduct.Title, result.Title)
	}

	// Other fields can be tested similarly
}
