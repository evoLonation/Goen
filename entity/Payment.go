package entity

type Payment struct {
	AmountTendered float64

	BelongedSale *Sale
}
