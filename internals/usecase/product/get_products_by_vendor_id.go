package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"sort"
)

func (u *productUseCase) GetProductsByVendorSortedByRating(vendorID int64, sortAscending bool) ([]*productEntities.Product, error) {
	ctx := context.Background()
	products, err := u.repo.GetAllByVendor(&ctx, vendorID)
	if err != nil {
		return nil, err
	}

	sort.Slice(products, func(i, j int) bool {
		if sortAscending {
			return products[i].Rating < products[j].Rating
		}
		return products[i].Rating > products[j].Rating
	})

	return products, nil
}

func (p *productUseCase) GetProductsByVendorGroupedByCategory(vendorID int64) (map[int][]*productEntities.Product, error) {
	ctx := context.Background()
	products, err := p.repo.GetAllByVendorGroupedByCategory(&ctx, vendorID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
