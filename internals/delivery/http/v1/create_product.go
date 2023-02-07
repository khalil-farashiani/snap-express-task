package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/khalil-farashiani/products-service/internals/dto"
	productUseCase "github.com/khalil-farashiani/products-service/internals/usecase/product"
	"net/http"
)

func createProduct(useCase productUseCase.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product dto.CreateProductRequest
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		res, err := useCase.CreateProduct(&ctx, product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}
