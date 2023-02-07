package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"github.com/khalil-farashiani/products-service/internals/dto"
)

type ProductUseCase interface {
	CreateProduct(*context.Context, dto.CreateProductRequest) (dto.CreateProductResponse, error)
	GetByID(ctx *context.Context, id int64) (*productEntities.Product, error)
	GetProductsByVendorSortedByRating(ctx *context.Context, vendorID int64, sortAscending bool) (dto.GetProductsByVendorResponse, error)
	GetProductsByVendorGroupedByCategory(ctx *context.Context, vendorID int64) (dto.GetProductsGroupedByCategoryResponse, error)
	PurchaseProduct(ctx *context.Context, productID int64) error
	Update(ctx *context.Context, id int64, updatedProduct *productEntities.Product) error
	GetNearbyProducts(ctx *context.Context, req dto.GetNearbyProductsRequest) (dto.GetNearByProductsResponse, error)
}

type productUseCase struct {
	repo productEntities.ProductRepository
}

func NewProductUseCase(repo productEntities.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}
