package entityForGen

type OrderStatus int

const (
	NEW       OrderStatus = 0
	RECEIVED  OrderStatus = 1
	REQUESTED OrderStatus = 2
)
