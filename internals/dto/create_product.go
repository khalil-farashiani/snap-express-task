package dto

import productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"

type CreateProductRequest struct {
	Title       string `json:"title"`
	TitleFa     string `json:"titleFa"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	VendorID    int    `json:"vendor_id"`
	CategoryID  int    `json:"category_id"`
	LocationID  int    `json:"location_id"`
	BrandID     int    `json:"brand_id"`
}

type CreateProductResponse struct {
	productEntities.Product
	Error string `json:"error,omitempty"`
}

func (c CreateProductRequest) CreateProductRequestToProductEntity() productEntities.Product {
	return productEntities.Product{
		Title:       c.Title,
		TitleFa:     c.TitleFa,
		Description: c.Description,
		Price:       c.Price,
		CategoryID:  c.CategoryID,
		VendorID:    c.VendorID,
		LocationID:  c.LocationID,
		BrandID:     c.BrandID,
		Stock:       c.Stock,
	}
}
