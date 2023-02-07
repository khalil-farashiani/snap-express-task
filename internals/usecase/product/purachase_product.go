package product

import (
	"context"
	"errors"
)

func (u *productUseCase) PurchaseProduct(productID int64) error {
	product, err := u.GetByID(productID)
	if err != nil {
		return err
	}
	if product.Stock <= 0 {
		return errors.New("product out of stock")
	}
	product.DecreaseStock(1)

	ctx := context.Background()
	return u.repo.Update(&ctx, product)
}
