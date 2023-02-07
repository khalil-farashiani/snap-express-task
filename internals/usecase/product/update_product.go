package product

import (
	"context"
	"github.com/khalil-farashiani/products-service/internals/dto"
)

func (p *productUseCase) Update(ctx *context.Context, id int64, updatedProduct *dto.UpdateProductRequest) error {
	currentProduct, err := p.GetByID(ctx, id)
	if err != nil {
		return err
	}
	currentProduct.Title = updatedProduct.Title
	currentProduct.TitleFa = updatedProduct.TitleFa
	currentProduct.BrandID = updatedProduct.BrandID
	currentProduct.CategoryID = updatedProduct.CategoryID
	currentProduct.Distance = updatedProduct.Distance
	currentProduct.Rating = updatedProduct.Rating
	currentProduct.Description = updatedProduct.Description
	currentProduct.Stock = updatedProduct.Stock
	currentProduct.Price = updatedProduct.Price

	return p.repo.Update(ctx, currentProduct)
}
