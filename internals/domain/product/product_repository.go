package product

import "context"

type ProductRepository interface {
	Store(*context.Context, *Product) error
	GetProductById(*context.Context, int64) (*Product, error)
}
