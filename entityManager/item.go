package entityManager

import entity2 "Cocome/entity"

type Item struct {
	entity
	Barcode     int      `db:"barcode"`
	Name        *string  `db:"name"`
	Price       *float64 `db:"price"`
	StockNumber *int     `db:"stock_number"`
	OrderPrice  *float64 `db:"order_price"`

	// other entity's * relation
	StoreId          *int `db:"store_id"`
	ProductCatalogId *int `db:"product_catalog_id"`
}

var ItemManager = &Manager[Item, *Item]{
	tableName: "item",
	idName:    "barcode",
}

func (p *Item) SetName(name string) {
	p.Name = &name
	p.addSetField("name")
}
func (p *Item) SetBarcode(barcode int) {
	p.Barcode = barcode
	p.addSetField("barcode")
}
func (p *Item) SetPrice(price float64) {
	p.Price = &price
	p.addSetField("price")
}
func (p *Item) SetOrderPrice(price float64) {
	p.OrderPrice = &price
	p.addSetField("order_price")
}

func (p *Item) SetStockNumber(stockNumber string) {
	p.Name = &stockNumber
	p.addSetField("stock_number")
}

func (p *Item) AddContainedSalesLine(salesLine *entity2.SalesLineItem) {
	// 让 salesLine的候选item键赋值为本对象
}
