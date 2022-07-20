package entityForGen

import "time"

type Sale struct {
	GoenId       int `gorm:"primaryKey"`
	Time         time.Time
	IsComplete   bool
	Amount       float64
	IsReadytoPay bool

	// reference
	// other entity's * relation
	StoreId       int
	CashDeskId    int
	CardPaymentId int
	CashPayment   int

	//BelongedStore         *Store
	//BelongedCashDesk      *CashDesk
	//ContainedSalesLine    []*SalesLineItem
	//AssoicatedCashPayment *CashPayment
	//AssoicatedCardPayment *CardPayment
}
