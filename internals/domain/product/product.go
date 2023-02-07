package product

type Product struct {
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
	Distance    float64
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
