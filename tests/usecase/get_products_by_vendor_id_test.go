package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"reflect"
	"testing"
)

func TestGetAllProductsByVendorID(t *testing.T) {
	// Define mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_product.NewMockProductRepository(ctrl)
	// Define expected products
	expectedProducts := []*productEntities.Product{
		{ID: 1, Title: "product 1", VendorID: 1, Rating: 4},
		{ID: 2, Title: "product 2", VendorID: 1, Rating: 5},
		{ID: 3, Title: "product 3", VendorID: 1, Rating: 3},
	}

	// Setup mock repository to return expected products

	// Create product use case with mock repository
	productUseCase := productUseCase.NewProductUseCase(m)

	// Call GetAllByVendorID method of product use case
	ctx := context.TODO()
	products, err := productUseCase.GetProductsByVendorSortedByRating(&ctx, 1, true)

	// Check if error occurs
	if err != nil {
		t.Errorf("Error while getting all products by vendor ID, got: %v", err)
	}

	// Check if the returned products are as expected
	if !reflect.DeepEqual(products, expectedProducts) {
		t.Errorf("Expected products: %v, but got: %v", expectedProducts, products)
	}

}

func TestGetProductsByVendorIDAndSortByRating(t *testing.T) {
	// Define mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_product.NewMockProductRepository(ctrl)

	//m.On("GetProductsByVendorIDAndSortByRating", 1, true).Return([]*productEntities.Product{
	//	{
	//		ID:          1,
	//		Title:       "product 1",
	//		Description: "product 1 description",
	//		Price:       100,
	//		Rating:      5,
	//		CategoryID:  1,
	//		VendorID:    1,
	//		LocationID:  1,
	//		BrandID:     1,
	//		Stock:       10,
	//	},
	//	{
	//		ID:          2,
	//		Title:       "product 2",
	//		Description: "product 2 description",
	//		Price:       200,
	//		Rating:      4,
	//		CategoryID:  2,
	//		VendorID:    1,
	//		LocationID:  2,
	//		BrandID:     2,
	//		Stock:       20,
	//	},
	//}, nil)

	productUseCase := productUseCase.NewProductUseCase(m)

	ctx := context.TODO()
	products, err := productUseCase.GetProductsByVendorSortedByRating(&ctx, 1, true)
	if err != nil {
		t.Errorf("Error while getting products, got: %v", err)
	}

	if len(products.Products) != 2 {
		t.Errorf("Expected to get 2 products, got: %d", len(products.Products))
	}

	if products.Products[0].ID != 1 {
		t.Errorf("Expected first product to have ID 1, got: %d", products.Products[0].ID)
	}

	if products.Products[1].ID != 2 {
		t.Errorf("Expected second product to have ID 2, got: %d", products.Products[1].ID)
	}
}

func TestGetProductsByVendorGroupedByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mock_product.NewMockProductRepository(ctrl)
	vendorID := 1
	products := []*productEntities.Product{
		{ID: 1, Title: "product 1", CategoryID: 1, VendorID: vendorID},
		{ID: 2, Title: "product 2", CategoryID: 2, VendorID: vendorID},
		{ID: 3, Title: "product 3", CategoryID: 1, VendorID: vendorID},
		{ID: 4, Title: "product 4", CategoryID: 3, VendorID: vendorID},
		{ID: 5, Title: "product 5", CategoryID: 1, VendorID: vendorID},
	}

	expectedResult := map[int][]*productEntities.Product{
		1: {products[0], products[2], products[4]},
		2: {products[1]},
		3: {products[3]},
	}

	ctx := context.TODO()
	mockProductRepository.EXPECT().GetAllByVendor(&ctx, vendorID).Return(products, nil)
	productUseCase := productUseCase.NewProductUseCase(mockProductRepository)
	result, err := productUseCase.GetProductsByVendorGroupedByCategory(&ctx, int64(vendorID))
	if err != nil {
		t.Errorf("Error while getting products, got: %v", err)
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}
