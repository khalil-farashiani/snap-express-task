package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

func (u *productUseCase) GetByID(id int64) (*productEntities.Product, error) {
	ctx := context.Background()
	p, err := u.repo.GetProductById(&ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
