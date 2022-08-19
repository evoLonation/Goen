package entityManager

type Item struct {
	Entity
	Barcode     int     `db:"barcode"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	StockNumber int     `db:"stock_number"`
	OrderPrice  float64 `db:"order_price"`

	BelongedItemGoenId *int `db:"belonged_item_goen_id"`
}

func (p *Item) SetName(name string) {
	p.Name = name
	p.addBasicField("name", name)
}
func (p *Item) SetBarcode(barcode int) {
	p.Barcode = barcode
	p.addBasicField("barcode", barcode)
}
func (p *Item) SetPrice(price float64) {
	p.Price = price
	p.addBasicField("price", price)
}
func (p *Item) SetOrderPrice(price float64) {
	p.OrderPrice = price
	p.addBasicField("order_price", price)
}

func (p *Item) SetStockNumber(stockNumber string) {
	p.Name = stockNumber
	p.addBasicField("stock_number", stockNumber)
}

func (p *Item) AddContainedItem(item *Item) {
	p.addJoinTableInsert(true, "contained_item", "item", item.GoenId)
}
func (p *Item) SetBelongedItem(item *Item) {
	p.BelongedItemGoenId = &item.GoenId
	p.addAssociationField("belonged_item_goen_id", item.GoenId)
}

func (p *Item) GetContainedItem() ([]*Item, error) {
	var items []*Item
	query := p.getSelectJoinTableQuery("contained_item", "item")
	if err := Db.Select(&items, query, p.GoenId); err != nil {
		return nil, err
	}
	for _, e := range items {
		e.initEntity(Founded, p.tableName, 0)
	}
	return items, nil
}
func (p *Item) GetBelongedItem() (*Item, error) {
	if p.BelongedItemGoenId == nil {
		return nil, nil
	} else {
		return ItemManager.Get(*p.BelongedItemGoenId)
	}
}
