package entity

import "time"

type CardPayment struct {
	Payment

	CardAccountNumber string
	ExpiryDate        time.Time

	//
	SaleId       int
	BelongedSale *Sale
}
