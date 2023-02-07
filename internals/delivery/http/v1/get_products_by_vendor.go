package v1

import (
	"github.com/gin-gonic/gin"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	"net/http"
	"strconv"
)

func GetProductsByVendorIdSortedByRating(useCase productUseCase.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		vendorId := c.Param("vendor_id")
		numericVendorId, err := strconv.Atoi(vendorId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		res, err := useCase.GetProductsByVendorSortedByRating(&ctx, int64(numericVendorId), true)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func GetProductsByVendorGroupedByCategory(useCase productUseCase.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		vendorId := c.Param("vendor_id")
		numericVendorId, err := strconv.Atoi(vendorId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		res, err := useCase.GetProductsByVendorGroupedByCategory(&ctx, int64(numericVendorId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}
