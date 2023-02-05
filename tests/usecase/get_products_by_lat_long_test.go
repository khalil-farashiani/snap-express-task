package usecase

import (
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
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
	sortByDistance := true
	products := []*productEntities.Product{
		{ID: 1, LocationID: 1},
		{ID: 2, LocationID: 2},
		{ID: 3, LocationID: 3},
	}
	mockProductRepo.EXPECT().GetAllByLocation(lat, long, sortByDistance).Return(products, nil)

	productUseCase := productUseCase.NewProductUseCase(mockProductRepo)
	gotProducts, err := productUseCase.GetNearbyProducts(lat, long, sortByDistance)
	if err != nil {
		t.Errorf("Error while getting nearby products, got: %v", err)
	}
	if len(gotProducts) != len(products) {
		t.Errorf("Number of nearby products is incorrect, want: %d, got: %d", len(products), len(gotProducts))
	}
}
