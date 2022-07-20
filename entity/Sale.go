package entity

import "time"

type Sale struct {
	Time         time.Time `gorm:"primaryKey"`
	IsComplete   bool
	Amount       float64
	IsReadytoPay bool

	// reference
	// other entity's * relation
	StoreId       int
	CashDeskId    int
	CardPaymentId int
	CashPayment   int

	BelongedStore         *Store
	BelongedCashDesk      *CashDesk
	ContainedSalesLine    []*SalesLineItem
	AssoicatedCashPayment *CashPayment
	AssoicatedCardPayment *CardPayment
}
