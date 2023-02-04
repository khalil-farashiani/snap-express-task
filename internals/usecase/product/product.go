package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
)

type ProductUseCase interface {
	CreateProduct(product *productEntities.Product) error
	GetByID(id int64) (*productEntities.Product, error)
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

func (u *productUseCase) GetByID(id int64) (*productEntities.Product, error) {
	ctx := context.Background()
	p, err := u.repo.GetProductById(&ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
