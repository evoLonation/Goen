package entityForGen

type SalesLineItem struct {
	GoenId    int `gorm:"primaryKey"`
	Quantity  int
	Subamount float64

	// other entity's * relation
	SaleId int
	ItemId int

	//BelongedSale *Sale
	//BelongedItem *Item
}
