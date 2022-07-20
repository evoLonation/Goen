package entityForGen

type Item struct {
	Barcode     int `gorm:"primaryKey"`
	Name        string
	Price       float64
	StockNumber int
	OrderPrice  float64

	// other entity's * relation
	StoreId          int
	ProductCatalogId int

	//BelongedCatalog *ProductCatalog
}
