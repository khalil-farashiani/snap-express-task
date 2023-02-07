package product

import (
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

type ProductUseCase interface {
	CreateProduct(product *productEntities.Product) error
	GetByID(id int64) (*productEntities.Product, error)
	GetProductsByVendorSortedByRating(vendorID int64, sortAscending bool) ([]*productEntities.Product, error)
	GetProductsByVendorGroupedByCategory(vendorID int64) (map[int][]*productEntities.Product, error)
	PurchaseProduct(productID int64) error
	Update(id int64, updatedProduct *productEntities.Product) error
	GetNearbyProducts(request GetNearbyProductsRequest) ([]*productEntities.Product, error)
}

type productUseCase struct {
	repo productEntities.ProductRepository
}

func NewProductUseCase(repo productEntities.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}
