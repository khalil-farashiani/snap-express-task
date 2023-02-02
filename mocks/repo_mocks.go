package mocks

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

type ProductRepository interface {
	Store() (ctx *context.Context, product productEntities.Product)
}
