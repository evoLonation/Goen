package entity

type Store struct {
	Id       int
	Name     string
	Address  string
	IsOpened bool

	//AssociationCashdeskes []*CashDesk
	//Productcatalogs       []*ProductCatalog
	//Items                 []*Item
	//Cashiers              []*Cashier
	//Sales                 []*Sale
}
