package domain

import (
	"testing"
)

func TestProduct(t *testing.T) {
	t.Run("Creating a new Product", func(t *testing.T) {
		id := 123
		faTitle := "(پودر قهوه اسپرسو الواتزا )250 گرم"
		description := "250 گرم"
		price := 259800
		rating := 5
		categoryID := 60
		vendorID := 12
		lat := 35.779052
		lon := 51.444655
		brandFa := "الواتزا"
		brandID := 379
		stock := 4

		product := NewProduct(id, faTitle, description, price, rating, categoryID, vendorID, lat, lon, brandFa, brandID, stock)

		if product.ID != id {
			t.Errorf("Expected product ID to be %d, but got %d", id, product.ID)
		}
		if product.FATitle != faTitle {
			t.Errorf("Expected product faTitle to be %s, but got %s", faTitle, product.FATitle)
		}
	})

	t.Run("Updating a Product's rating", func(t *testing.T) {
		// Given
		product := NewProduct(...)
		newRating := 4

		// When
		product.UpdateRating(newRating)

		// Then
		if product.Rating != newRating {
			t.Errorf("Expected product rating to be %d, but got %d", newRating, product.Rating)
		}
	})

	// TODO: Add more tests here...
}

