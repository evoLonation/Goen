package entity

import (
	"Cocome/entityManager"
)

var paymentManager entityManager.ManagerForEntity[Payment]
var PaymentManager entityManager.InheritManagerForOther[Payment]

type Payment interface {
	SetAmountTendered(amountTendered float64)
	GetAmountTendered() float64
}

type PaymentEntity struct {
	entityManager.BasicEntity

	AmountTendered float64 `db:"amount_tendered"`
}

func (p *PaymentEntity) SetAmountTendered(amountTendered float64) {
	p.AmountTendered = amountTendered
	p.AddBasicFieldChange("amount_tendered")
}

func (p *PaymentEntity) GetAmountTendered() float64 {
	return p.AmountTendered
}
