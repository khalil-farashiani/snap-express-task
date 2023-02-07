package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/khalil-farashiani/products-service/internals/domain/location"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"github.com/khalil-farashiani/products-service/internals/dto"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestGetNearbyProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mock_product.NewMockProductRepository(ctrl)

	lat := 1.23
	long := 4.56
	products := []*productEntities.Product{
		{ID: 1, LocationID: 1},
		{ID: 2, LocationID: 2},
		{ID: 3, LocationID: 3},
	}
	ctx := context.TODO()
	mockProductRepo.EXPECT().GetAllByLocation(&ctx, location.Location{Lat: 0, Lon: 0}, true).Return(products, nil)

	req := dto.GetNearbyProductsRequest{
		Latitude:   lat,
		Longitude:  long,
		SortOption: "distance",
	}
	productUseCase := productUseCase.NewProductUseCase(mockProductRepo)
	gotProducts, err := productUseCase.GetNearbyProducts(&ctx, req)
	if err != nil {
		t.Errorf("Error while getting nearby products, got: %v", err)
	}
	if len(gotProducts.Products) != len(products) {
		t.Errorf("Number of nearby products is incorrect, want: %d, got: %d", len(products), len(gotProducts.Products))
	}
}
