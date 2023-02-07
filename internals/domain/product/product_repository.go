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
	Update(*context.Context, *Product) error
	GetAllByVendorGroupedByCategory(*context.Context, int64) (map[int][]*Product, error)
	GetAll(*context.Context) ([]*Product, error)
	GetLocationByID(*context.Context, int64) (location.Location, error)
}
