package dto

import productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"

type UpdateProductRequest struct {
	productEntities.Product
}

type UpdateProductResponse struct {
	productEntities.Product
}
