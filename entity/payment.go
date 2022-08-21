package entity

import (
	"Cocome/entityManager"
)

var paymentManager entityManager.ManagerForEntity[Payment]
var PaymentManager entityManager.ManagerForOther[Payment]

func init() {

}

type PaymentGetSet interface {
	SetAmountTendered(amountTendered float64)
	GetAmountTendered() float64
}

type Payment interface {
	PaymentGetSet
	GetRealType() entityManager.GoenInheritType
	TurnToCardPayment() (CardPayment, error)
}

type PaymentEntity struct {
	entityManager.BasicEntity

	AmountTendered float64 `db:"amount_tendered"`
}

func (p *PaymentEntity) TurnToCardPayment() (CardPayment, error) {
	return cardPaymentManager.Get(p.GoenId)
}

func (p *PaymentEntity) SetAmountTendered(amountTendered float64) {
	p.AmountTendered = amountTendered
	p.AddBasicFieldChange("amount_tendered")
}

func (p *PaymentEntity) GetAmountTendered() float64 {
	return p.AmountTendered
}
