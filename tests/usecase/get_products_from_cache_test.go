package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"reflect"
	"testing"
)

func TestCacheGetProductsByVendor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	vendorID := 1
	products := []*productEntities.Product{
		{ID: 1, Title: "product 1", VendorID: vendorID, Stock: 10},
		{ID: 2, Title: "product 2", VendorID: vendorID, Stock: 0},
		{ID: 3, Title: "product 3", VendorID: vendorID, Stock: 5},
	}

	expectedResult := []*productEntities.Product{
		products[0], products[2],
	}

	mockProductRepository := mock_product.NewMockProductRepository(ctrl)

	ctx := context.TODO()
	mockProductRepository.EXPECT().GetAllByVendor(&ctx, vendorID).Return(products, nil).Times(1)

	cache := productUseCase.NewCache(mockProductRepository)
	result, err := cache.GetProductsByVendor(vendorID)
	if err != nil {
		t.Errorf("Error while getting products, got: %v", err)
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}

	// make sure the repository is not called again
	result, err = cache.GetProductsByVendor(vendorID)
	if err != nil {
		t.Errorf("Error while getting products, got: %v", err)
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}
