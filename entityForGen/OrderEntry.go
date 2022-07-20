package entity

type OrderEntry struct {
	Quantity  int
	SubAmount float64

	//
	OrderProductId int

	ItemId int
	//Item   *Item
}
