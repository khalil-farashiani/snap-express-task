package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

func (p *productUseCase) Update(id int64, updatedProduct *productEntities.Product) error {
	currentProduct, err := p.GetByID(id)
	if err != nil {
		return err
	}
	currentProduct.Title = updatedProduct.Title
	currentProduct.Description = updatedProduct.Description
	currentProduct.Stock = updatedProduct.Stock
	currentProduct.Price = updatedProduct.Price

	ctx := context.Background()
	return p.repo.Update(&ctx, currentProduct)
}
