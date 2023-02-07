package v1

import (
	"github.com/gin-gonic/gin"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	"net/http"
	"strconv"
)

func PurchaseProduct(useCase productUseCase.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("product_id")
		numericProductId, err := strconv.Atoi(productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		err = useCase.PurchaseProduct(&ctx, int64(numericProductId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "successfully purchased"})
	}
}
