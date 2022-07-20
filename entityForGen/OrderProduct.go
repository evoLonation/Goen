package entityForGen

import "time"

type OrderProduct struct {
	Id          int
	Time        time.Time
	OrderStatus OrderStatus
	Amount      float64

	// reference
	SupplierId int
	//Supplier   *Supplier

	//ContainedEntries []*OrderEntry
}
