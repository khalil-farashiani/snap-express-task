package product

type Product struct {
	ID          int     `bson:"_id"`
	Title       string  `bson:"title"`
	TitleFa     string  `bson:"title_fa"`
	Description string  `bson:"description"`
	Price       int     `bson:"price"`
	Rating      int     `bson:"rating"`
	CategoryID  int     `bson:"category_id"`
	VendorID    int     `bson:"vendor_id"`
	LocationID  int     `bson:"location_id"`
	BrandID     int     `bson:"brand_id"`
	Stock       int     `bson:"stock"`
	Distance    float64 `bson:"distance"`
}

func (p *Product) IncreaseStock(count int) {
	p.Stock += count
}

func (p *Product) DecreaseStock(count int) {
	p.Stock -= count
}

func (p *Product) UpdateRating(rating int) {
	p.Rating = rating
}

func NewProduct(id int, titleFa string, titleEn string, description string, price int, rating int, categoryID int, vendorID int, locationId int, brandID int, stock int) *Product {
	return &Product{
		ID:          id,
		TitleFa:     titleFa,
		Title:       titleEn,
		Description: description,
		Price:       price,
		Rating:      rating,
		CategoryID:  categoryID,
		VendorID:    vendorID,
		LocationID:  locationId,
		BrandID:     brandID,
		Stock:       stock,
	}
}
