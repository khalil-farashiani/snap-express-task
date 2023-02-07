package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

func (u *productUseCase) CreateProduct(product *productEntities.Product) error {
	ctx := context.Background()
	err := u.repo.Store(&ctx, product)
	if err != nil {
		return err
	}
	return nil
}
