package entity

type CashPayment struct {
	Payment

	Balance float64

	//
	SaleId int
	//BelongedSale *Sale
}
