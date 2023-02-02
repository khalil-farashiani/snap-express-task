package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

type ProductUseCase interface {
	CreateProduct(product *productEntities.Product) error
}

type productUseCase struct {
	repo productEntities.ProductRepository
}

func NewProductUseCase(repo productEntities.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}

func (u *productUseCase) CreateProduct(product *productEntities.Product) error {
	ctx := context.Background()
	err := u.repo.Store(&ctx, product)
	if err != nil {
		return err
	}
	return nil
}
