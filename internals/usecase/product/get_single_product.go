package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

func (p *productUseCase) GetByID(ctx *context.Context, id int64) (*productEntities.Product, error) {
	products, err := p.repo.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
