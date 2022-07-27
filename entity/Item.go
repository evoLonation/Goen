package entity

type Item struct {
	Barcode     int     `db:"barcode"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	StockNumber int     `db:"stock_number"`
	OrderPrice  float64 `db:"order_price"`

	// other entity's * relation
	StoreId          int `db:"store_id"`
	ProductCatalogId int `db:"product_catalog_id"`

	BelongedCatalog *ProductCatalog
}
