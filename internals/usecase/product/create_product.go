package product

import (
	"context"
	"github.com/khalil-farashiani/products-service/internals/dto"
)

func (p *productUseCase) CreateProduct(ctx *context.Context, req dto.CreateProductRequest) (dto.CreateProductResponse, error) {
	product := req.CreateProductRequestToProductEntity()
	err := p.repo.Store(ctx, &product)
	if err != nil {
		return dto.CreateProductResponse{}, err
	}
	return dto.CreateProductResponse{
		Product: product,
		Error:   "",
	}, nil
}
