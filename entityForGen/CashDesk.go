package entityForGen

type CashDesk struct {
	Id       int
	Name     string
	IsOpened bool

	// reference
	// other entity's * relation
	StoreId int

	//ContainedSales []*Sale
	//BelongedStore  *Store
}
