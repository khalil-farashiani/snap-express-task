package product

import (
	"context"
	"github.com/khalil-farashiani/products-service/internals/dto"
	"sort"
)

func (p *productUseCase) GetProductsByVendorSortedByRating(ctx *context.Context, vendorID int64, sortAscending bool) (dto.GetProductsByVendorResponse, error) {
	products, err := p.repo.GetAllByVendor(ctx, vendorID)
	if err != nil {
		return dto.GetProductsByVendorResponse{}, err
	}

	sort.Slice(products, func(i, j int) bool {
		if sortAscending {
			return products[i].Rating < products[j].Rating
		}
		return products[i].Rating > products[j].Rating
	})

	return dto.GetProductsByVendorResponse{Products: products}, nil
}

func (p *productUseCase) GetProductsByVendorGroupedByCategory(ctx *context.Context, vendorID int64) (dto.GetProductsGroupedByCategoryResponse, error) {
	products, err := p.repo.GetAllByVendorGroupedByCategory(ctx, vendorID)
	if err != nil {
		return dto.GetProductsGroupedByCategoryResponse{}, err
	}
	return dto.GetProductsGroupedByCategoryResponse{Products: products}, nil
}
