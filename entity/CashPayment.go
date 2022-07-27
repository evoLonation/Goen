package entity

import "database/sql"

type CashPayment struct {
	Payment

	Balance sql.NullFloat64

	//
	SaleId       int
	BelongedSale *Sale
}
