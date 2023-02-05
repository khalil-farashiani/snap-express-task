package usecase

import (
	"github.com/golang/mock/gomock"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	mock_product "github.com/khalil-farashiani/products-service/mocks"
	"testing"
)

func TestChangeStockForMultipleProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mock_product.NewMockProductRepository(ctrl)
	productUseCase := productUseCase.NewProductUseCase(mockProductRepository)

	productIDs := []int{1, 2, 3}
	newStockValues := []int{10, 20, 30}

	mockProductRepository.EXPECT().UpdateStockForMultipleProducts(productIDs, newStockValues).Return(nil)

	err := productUseCase.ChangeStockForMultipleProducts(productIDs, newStockValues)
	if err != nil {
		t.Errorf("Error while changing stock for multiple products, got: %v", err)
	}
}
