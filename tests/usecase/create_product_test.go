package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_product.NewMockProductRepository(ctrl)
	productUseCase := productUseCase.NewProductUseCase(m)

	//mock the database
	ctx := context.Background()
	m.EXPECT().Store(&ctx, product).DoAndReturn(func(ctx *context.Context, p *productEntities.Product) error {
		p.ID = 1
		return nil
	}).AnyTimes()

	err := productUseCase.CreateProduct(product)
	if err != nil {
		t.Errorf("Error while creating product, got: %v", err)
	}
	if product.ID == 0 {
		t.Errorf("Product ID is not set, got: %d", product.ID)
	}
}
