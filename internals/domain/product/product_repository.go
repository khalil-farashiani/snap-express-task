package product

import (
	"context"
	"github.com/khalil-farashiani/products-service/internals/domain/location"
)

type ProductRepository interface {
	Store(*context.Context, *Product) error
	BulkStore(*context.Context, []*Product) error
	GetProductById(*context.Context, int64) (*Product, error)
	UpdateProduct(*context.Context, *Product) error
	GetAllByVendor(*context.Context, int64) ([]*Product, error)
	GetProductsByVendorIDAndSortByRating(*context.Context, int64) ([]*Product, error)
	GetAllByLocation(*context.Context, location.Location, bool) ([]*Product, error)
	UpdateStockForMultipleProducts(*context.Context, []int64, []int64) error
}
