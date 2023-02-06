package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestPurchaseProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := &productEntities.Product{
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

	mockProductRepository := mock_product.NewMockProductRepository(ctrl)
	ctx := context.TODO()
	mockProductRepository.EXPECT().GetProductById(&ctx, gomock.Eq(product.ID)).Return(product, nil)
	mockProductRepository.EXPECT().UpdateProduct(&ctx, gomock.Eq(product)).Return(nil)

	productUseCase := productUseCase.NewProductUseCase(mockProductRepository)
	err := productUseCase.PurchaseProduct(product.ID)
	if err != nil {
		t.Errorf("Error while purchasing product, got: %v", err)
	}
	if product.Stock != 9 {
		t.Errorf("Stock of product is not decreased by 1, got: %d", product.Stock)
	}
}
