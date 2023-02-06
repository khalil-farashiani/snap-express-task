package usecase

import (
	"fmt"
	"github.com/golang/mock/gomock"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"os"
	"testing"
)

func TestCreateMultipleProductsFromFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mock_product.NewMockProductRepository(ctrl)
	products := []*productEntities.Product{
		{ID: 1, Title: "product 1", CategoryID: 1, VendorID: 1, Stock: 10},
		{ID: 2, Title: "product 2", CategoryID: 2, VendorID: 2, Stock: 20},
		{ID: 3, Title: "product 3", CategoryID: 3, VendorID: 3, Stock: 30},
		{ID: 4, Title: "product 4", CategoryID: 4, VendorID: 4, Stock: 40},
	}

	file, err := os.Create("test_products.txt")
	if err != nil {
		t.Errorf("Error while creating file, got: %v", err)
	}
	defer os.Remove("test_products.txt")

	for _, product := range products {
		_, err := file.WriteString(fmt.Sprintf("%d,%s,%d,%d,%d\n", product.ID, product.Title, product.CategoryID, product.VendorID, product.Stock))
		if err != nil {
			t.Errorf("Error while writing to file, got: %v", err)
		}
	}

	mockProductRepository.EXPECT().BulkStore(gomock.Any()).Times(len(products)).Return(nil)

	productUseCase := productUseCase.NewProductUseCase(mockProductRepository)
	err = productUseCase.CreateMultipleProductsFromFile("test_products.txt")
	if err != nil {
		t.Errorf("Error while creating multiple products, got: %v", err)
	}
}
