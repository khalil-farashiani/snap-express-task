package product

import "context"

type ProductRepository interface {
	Store(*context.Context, *Product) error
}
