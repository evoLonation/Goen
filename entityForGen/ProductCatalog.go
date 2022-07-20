package entity

type ProductCatalog struct {
	Id   int
	Name string

	// other entity's * relation
	StoreId int

	//ContainedItems []*Item
}
