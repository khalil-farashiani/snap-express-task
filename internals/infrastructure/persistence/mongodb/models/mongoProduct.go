package models

import (
	"context"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoProduct is a struct that will represent a Product in the MongoDB database
type MongoProduct struct {
	ID          int
	Title       string
	TitleFa     string
	Description string
	Price       int
	Rating      int
	CategoryID  int
	VendorID    int
	LocationID  int
	BrandID     int
	Stock       int
}

// ProductRepository is a struct that implements the ProductRepository interface
type ProductRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewProductRepository returns a new instance of ProductRepository
func NewProductRepository(connString string) (*ProductRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	db := client.Database("products")
	return &ProductRepository{client: client, db: db}, nil
}

// Store stores a new product in the MongoDB database
func (r *ProductRepository) Store(ctx *context.Context, p *productEntities.Product) error {
	collection := r.db.Collection("products")
	_, err := collection.InsertOne(*ctx, &MongoProduct{
		ID:          p.ID,
		Title:       p.Title,
		TitleFa:     p.TitleFa,
		Description: p.Description,
		Price:       p.Price,
		Rating:      p.Rating,
		CategoryID:  p.CategoryID,
		VendorID:    p.VendorID,
		LocationID:  p.LocationID,
		BrandID:     p.BrandID,
		Stock:       p.Stock,
	})
	if err != nil {
		log.Println("Error inserting product into database:", err)
		return err
	}

	return nil
}
