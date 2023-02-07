package dto

import productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"

type GetNearbyProductsRequest struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	SortOption string  `json:"sortOption"`
}

type GetNearByProductsResponse struct {
	Products []*productEntities.Product
}
