package entityManager

//
//type SalesLineItem struct {
//	GoenId    int     `db:"goen_id"`
//	Quantity  int     `db:"quantity"`
//	Subamount float64 `db:"subamount"`
//
//	// other Entity's * relation
//	GoenItemContainedSalesLine int `db:"goen_item_contained_sales_line"`
//	ItemId                     int `db:"item_id"`
//	SaleId                     int `db:"sale_id"`
//
//	//BelongedSale *Sale
//	//BelongedItem *Item
//}
//
//var SalesLineItemManager = ManagerGeneric[SalesLineItem, *SalesLineItem]{
//	tableName: "sales_line_item",
//	idName:    "goen_id",
//}
