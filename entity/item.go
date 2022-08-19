package entity

import "Cocome/entityManager"

var ItemManager *entityManager.ManagerGeneric[Item, *Item] = entityManager.NewManager[Item, *Item]("item")

type Item struct {
	entityManager.Entity
	Barcode     int     `db:"barcode"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	StockNumber int     `db:"stock_number"`
	OrderPrice  float64 `db:"order_price"`

	BelongedItemGoenId *int `db:"belonged_item_goen_id"`
}

func (p *Item) SetName(name string) {
	p.Name = name
	p.BasicFieldChange("name")
}
func (p *Item) SetBarcode(barcode int) {
	p.Barcode = barcode
	p.BasicFieldChange("barcode")
}
func (p *Item) SetPrice(price float64) {
	p.Price = price
	p.BasicFieldChange("price")
}
func (p *Item) SetOrderPrice(price float64) {
	p.OrderPrice = price
	p.BasicFieldChange("order_price")
}

func (p *Item) SetStockNumber(stockNumber string) {
	p.Name = stockNumber
	p.BasicFieldChange("stock_number")
}

func (p *Item) AddContainedItem(item *Item) {
	p.MultiAssChange(entityManager.Include, "item_contained_item", item.GoenId)
}
func (p *Item) SetBelongedItem(item *Item) {
	p.BelongedItemGoenId = &item.GoenId
	p.AssFieldChange("belonged_item_goen_id")
}

func (p *Item) GetContainedItem() ([]*Item, error) {
	return ItemManager.FindFromMultiAssTable("item_contained_item", p.GoenId)
}

func (p *Item) GetBelongedItem() (*Item, error) {
	if p.BelongedItemGoenId == nil {
		return nil, nil
	} else {
		return ItemManager.Get(*p.BelongedItemGoenId)
	}
}
