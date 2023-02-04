package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestGetProductByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_product.NewMockProductRepository(ctrl)

	p := &productEntities.Product{
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
	//mock the database
	ctx := context.Background()
	m.EXPECT().GetProductById(&ctx, 1).DoAndReturn(func(ctx *context.Context, id int64) *productEntities.Product {
		return p
	}).AnyTimes()
	productUseCase := productUseCase.NewProductUseCase(m)

	result, err := productUseCase.GetByID(1)
	if err != nil {
		t.Errorf("Error while getting product, got: %v", err)
	}

	if result == nil {
		t.Errorf("No product found, got: %v", result)
	}

	if result.ID != p.ID {
		t.Errorf("Product ID does not match, expected: %d, got: %d", mockProduct.ID, result.ID)
	}

	if result.Title != p.Title {
		t.Errorf("Product title does not match, expected: %s, got: %s", mockProduct.Title, result.Title)
	}

	// Other fields can be tested similarly
}
