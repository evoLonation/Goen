package entity

import (
	"database/sql"
)

type Sale struct {
	GoenId       int             `db:"GoenId"`
	Time         sql.NullTime    `db:"Time"`
	IsComplete   bool            `db:"IsComplete"`
	Amount       sql.NullFloat64 `db:"Amount"`
	IsReadytoPay bool            `db:"IsReadytoPay"`

	// reference
	// other entity's * relation
	StoreId       int
	CashDeskId    sql.NullInt32 `db:"CashDeskId"`
	CardPaymentId int
	CashPayment   int

	BelongedStore         *Store
	BelongedCashDesk      *CashDesk
	ContainedSalesLine    []*SalesLineItem
	AssoicatedCashPayment *CashPayment
	AssoicatedCardPayment *CardPayment
}
