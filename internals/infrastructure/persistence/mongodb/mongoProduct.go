package mongodb

import (
	"context"
	"github.com/khalil-farashiani/products-service/internals/domain/location"
	productEntities "github.com/khalil-farashiani/products-service/internals/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName                = "snap_express"
	productCollectionName = "products"
)

// ProductRepository is a struct that implements the ProductRepository interface
type ProductRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
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

	db := client.Database(dbName)
	return &ProductRepository{client: client, db: db}, nil
}

// Store stores a new product in the MongoDB database
func (r *ProductRepository) Store(ctx *context.Context, p *productEntities.Product) error {
	collection := r.db.Collection(productCollectionName)
	_, err := collection.InsertOne(*ctx, &productEntities.Product{
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

func (m *ProductRepository) BulkStore(ctx *context.Context, products []*productEntities.Product) error {
	docs := make([]interface{}, len(products))
	for i, product := range products {
		docs[i] = product
	}
	_, err := m.collection.InsertMany(*ctx, docs)
	return err
}

func (m *ProductRepository) GetProductById(ctx *context.Context, id int64) (*productEntities.Product, error) {
	filter := bson.M{"id": id}
	var product productEntities.Product
	err := m.collection.FindOne(*ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m *ProductRepository) GetAllByVendor(ctx *context.Context, vendorID int64) ([]*productEntities.Product, error) {
	filter := bson.M{"vendor_id": vendorID}
	cur, err := m.collection.Find(*ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)
	var products []*productEntities.Product
	for cur.Next(*ctx) {
		var product productEntities.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (ms *ProductRepository) GetProductsByVendorIDAndSortByRating(ctx *context.Context, vendorID int64) ([]*productEntities.Product, error) {
	var products []*productEntities.Product
	filter := bson.M{"vendor_id": vendorID}
	cur, err := ms.db.Collection(productCollectionName).Find(*ctx, filter, options.Find().SetSort(bson.M{"rating": -1}))
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)
	for cur.Next(*ctx) {
		var product productEntities.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (ms *ProductRepository) GetAllByLocation(ctx *context.Context, location location.Location, availableOnly bool) ([]*productEntities.Product, error) {
	var products []*productEntities.Product
	filter := bson.M{"location_id": location}
	if availableOnly {
		filter["stock"] = bson.M{"$gt": 0}
	}
	cur, err := ms.db.Collection(productCollectionName).Find(*ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)
	for cur.Next(*ctx) {
		var product productEntities.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (ms *ProductRepository) UpdateStockForMultipleProducts(ctx *context.Context, productIDs []int64, stocks []int64) error {
	bulkOps := make([]mongo.WriteModel, len(productIDs))
	for i := range productIDs {
		bulkOps[i] = mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": productIDs[i]}).
			SetUpdate(bson.M{"$set": bson.M{"stock": stocks[i]}})
	}
	_, err := ms.db.Collection(productCollectionName).BulkWrite(*ctx, bulkOps)
	return err
}

func (ms *ProductRepository) Update(ctx *context.Context, product *productEntities.Product) error {
	_, err := ms.db.Collection(productCollectionName).UpdateOne(*ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	return err
}

func (m *ProductRepository) GetAllByVendorGroupedByCategory(ctx *context.Context, vendorID int64) (map[int][]*productEntities.Product, error) {
	// Get a handle to the products collection
	collection := m.client.Database("snap_express").Collection("product")

	// Create a pipeline to match the vendorID and group by category
	pipeline := []bson.M{
		{"$match": bson.M{"vendorID": vendorID}},
		{"$group": bson.M{"_id": "$category", "products": bson.M{"$push": "$$ROOT"}}},
	}

	// Execute the pipeline and get the results
	cur, err := collection.Aggregate(*ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)

	// Create a map to store the results
	result := make(map[int][]*productEntities.Product)

	// Iterate through the results and add them to the map
	for cur.Next(*ctx) {
		var group bson.M
		if err := cur.Decode(&group); err != nil {
			return nil, err
		}

		var products []*productEntities.Product
		var productBytes [][]byte

		for _, item := range group["products"].(bson.A) {
			b, err := bson.Marshal(item.(bson.M))
			if err != nil {
				return nil, err
			}
			productBytes = append(productBytes, b)
		}
		for _, b := range productBytes {
			var product productEntities.Product
			err := bson.Unmarshal(b, &product)
			if err != nil {
				return nil, err
			}
			products = append(products, &product)
		}

		result[int(group["_id"].(int32))] = products
	}

	return result, nil
}

func (p *ProductRepository) GetAll(ctx *context.Context) ([]*productEntities.Product, error) {
	var products []*productEntities.Product
	collection := p.client.Database(dbName).Collection(productCollectionName)

	cursor, err := collection.Find(*ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(*ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetLocationByID(ctx *context.Context, locationID int64) (location.Location, error) {
	var location location.Location
	collection := r.client.Database("store").Collection("locations")

	err := collection.FindOne(*ctx, bson.M{"_id": locationID}).Decode(&location)
	if err != nil {
		return location, err
	}

	return location, nil
}

func (r *ProductRepository) UpdateProduct(ctx *context.Context, product *productEntities.Product) error {
	collection := r.client.Database("dbName").Collection("productCollection")
	_, err := collection.UpdateOne(*ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}
