package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"math"
	"sort"
)

const (
	RadiusOfEarth = 6371 // Radius of Earth in kilometers
)

type GetNearbyProductsRequest struct {
	Latitude   float64
	Longitude  float64
	SortOption string
}

func (u *productUseCase) GetNearbyProducts(request GetNearbyProductsRequest) ([]*productEntities.Product, error) {
	ctx := context.Background()
	products, err := u.repo.GetAll(&ctx)
	if err != nil {
		return nil, err
	}

	var nearbyProducts = make([]*productEntities.Product, 100)
	for _, product := range products {
		location, err := u.repo.GetLocationByID(&ctx, int64(product.LocationID))
		if err != nil {
			return nil, err
		}

		distance := calculateDistance(request.Latitude, request.Longitude, location.Lat, location.Lon)
		product.Distance = distance
		nearbyProducts = append(nearbyProducts, product)
	}

	if request.SortOption == "distance" {
		sort.Slice(nearbyProducts, func(i, j int) bool {
			return nearbyProducts[i].Distance < nearbyProducts[j].Distance
		})
	}

	return nearbyProducts, nil
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	sLat1 := lat1 * (math.Pi / 180)
	sLat2 := lat2 * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(sLat1)*math.Cos(sLat2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return RadiusOfEarth * c
}
