package domain

import (
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"testing"
)

const (
	id          = 123
	enTitle     = "coffee"
	faTitle     = "(پودر قهوه اسپرسو الواتزا )250 گرم"
	description = "250 گرم"
	price       = 259800
	rating      = 5
	categoryID  = 60
	vendorID    = 12
	locationId  = 2
	brandID     = 379
	stock       = 4
)

func TestCreateProduct(t *testing.T) {
	t.Run("Creating a new Product", func(t *testing.T) {
		product := productEntities.NewProduct(id, faTitle, enTitle, description, price, rating, categoryID, vendorID, locationId, brandID, stock)

		if product.ID != id {
			t.Errorf("Expected product ID to be %d, but got %d", id, product.ID)
		}
		if product.TitleFa != faTitle {
			t.Errorf("Expected product faTitle to be %s, but got %s", faTitle, product.TitleFa)
		}
	})
}

func TestUpdateRating(t *testing.T) {
	t.Run("Updating a Product's rating", func(t *testing.T) {
		product := productEntities.NewProduct(id, faTitle, enTitle, description, price, rating, categoryID, vendorID, locationId, brandID, stock)
		newRating := 4

		product.UpdateRating(newRating)

		if product.Rating != newRating {
			t.Errorf("Expected product rating to be %d, but got %d", newRating, product.Rating)
		}
	})
}

func TestIncreaseStock(t *testing.T) {
	t.Run("Increase a Product's Stock", func(t *testing.T) {
		product := productEntities.NewProduct(id, faTitle, enTitle, description, price, rating, categoryID, vendorID, locationId, brandID, stock)
		stockCount := 4

		product.IncreaseStock(stockCount)

		if product.Stock != stockCount+stock {
			t.Errorf("Expected product stock to be %d, but got %d", stockCount+stock, product.Stock)
		}
	})
}

func TestDecreaseStock(t *testing.T) {
	t.Run("decrease a Product's Stock", func(t *testing.T) {
		product := productEntities.NewProduct(id, faTitle, enTitle, description, price, rating, categoryID, vendorID, locationId, brandID, stock)
		stockCount := 3

		product.DecreaseStock(stockCount)

		if product.Stock != stock-stockCount {
			t.Errorf("Expected product stock to be %d, but got %d", stock-stockCount, product.Stock)
		}
	})
}
