package product

import (
	"context"
	"errors"
)

func (p *productUseCase) PurchaseProduct(ctx *context.Context, productID int64) error {
	product, err := p.GetByID(ctx, productID)
	if err != nil {
		return err
	}
	if product.Stock <= 0 {
		return errors.New("product out of stock")
	}
	product.DecreaseStock(1)

	return p.repo.Update(ctx, product)
}
