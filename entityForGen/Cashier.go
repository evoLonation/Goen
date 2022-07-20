package entityForGen

type Cashier struct {
	Id   int
	Name string

	// other entity's * relation
	StoreId int

	//WorkedStore *Store
}
