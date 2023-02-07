package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/khalil-farashiani/products-service/internals/delivery/http/v1"
	"github.com/khalil-farashiani/products-service/internals/usecase/product"
)

func GetRouts(r *gin.Engine, store product.ProductUseCase) {
	r.POST("CreateProduct", v1.CreateProduct(store))
	r.POST("GetNearByProducts", v1.GetNearByProducts(store))
	r.GET("GetProductByID/:product_id", v1.GetProductByID(store))
	r.GET("GetProductsByVendorIdSortedByRating/:vendor_id", v1.GetProductsByVendorIdSortedByRating(store))
	r.GET("GetProductsByVendorGroupedByCategory/:vendor_id", v1.GetProductsByVendorGroupedByCategory(store))
	r.GET("PurchaseProduct:/product_id", v1.PurchaseProduct(store))
	r.PATCH("UpdateProduct", v1.UpdateProduct(store))
}
