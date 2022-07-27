package entity

type SalesLineItem struct {
	GoenId    int
	Quantity  int
	Subamount float64

	// other entity's * relation
	SaleId int `db:"sale_id"`
	ItemId int `db:"item_id"`

	BelongedSale *Sale
	BelongedItem *Item
}
