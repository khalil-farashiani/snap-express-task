package dto

import productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"

type GetProductsGroupedByCategoryResponse struct {
	Products map[int][]*productEntities.Product
}

type GetProductsByVendorResponse struct {
	Products []*productEntities.Product
}
