package product

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"github.com/khalil-farashiani/products-service/internals/dto"
	"math"
	"sort"
)

const (
	RadiusOfEarth = 6371 // Radius of Earth in kilometers
)

func (p *productUseCase) GetNearbyProducts(ctx *context.Context, req dto.GetNearbyProductsRequest) (dto.GetNearByProductsResponse, error) {
	products, err := p.repo.GetAll(ctx)
	if err != nil {
		return dto.GetNearByProductsResponse{}, err
	}

	var nearbyProducts = make([]*productEntities.Product, 100)
	for _, product := range products {
		location, err := p.repo.GetLocationByID(ctx, int64(product.LocationID))
		if err != nil {
			return dto.GetNearByProductsResponse{}, err
		}

		distance := calculateDistance(req, location.Lat, location.Lon)
		product.Distance = distance
		nearbyProducts = append(nearbyProducts, product)
	}

	if req.SortOption == "distance" {
		sort.Slice(nearbyProducts, func(i, j int) bool {
			return nearbyProducts[i].Distance < nearbyProducts[j].Distance
		})
	}

	return dto.GetNearByProductsResponse{Products: nearbyProducts}, nil
}

func calculateDistance(req dto.GetNearbyProductsRequest, lat2, lon2 float64) float64 {
	dLat := (lat2 - req.Latitude) * (math.Pi / 180)
	dLon := (lon2 - req.Longitude) * (math.Pi / 180)

	sLat1 := req.Latitude * (math.Pi / 180)
	sLat2 := lat2 * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(sLat1)*math.Cos(sLat2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return RadiusOfEarth * c
}
