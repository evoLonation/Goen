package entity

import (
	"Cocome/entityManager"
	"time"
)

const (
	CardPaymentInheritType entityManager.GoenInheritType = 1
)

var paymentManager entityManager.ManagerForEntity[*Payment] = entityManager.NewManager[Payment, *Payment]("payment")
var PaymentManager entityManager.ManagerForOther[*Payment] = paymentManager.(entityManager.ManagerForOther[*Payment])
var cardPaymentManager entityManager.ManagerForEntity[*CardPayment] = entityManager.NewInheritManager[CardPayment, *CardPayment]("card_payment", PaymentManager.(entityManager.InheritManagerForRecur), CardPaymentInheritType)
var CardPaymentManager entityManager.ManagerForOther[*CardPayment] = cardPaymentManager.(entityManager.ManagerForOther[*CardPayment])

type PaymentInterface interface {
	GetRealType() entityManager.GoenInheritType
	SetAmountTendered(amountTendered float64)
	GetAmountTendered() float64
	GetGoenId() int
	TurnToCardPayment() (*CardPayment, error)
}

type Payment struct {
	entityManager.BasicEntity
	AmountTendered float64 `db:"amount_tendered"`
}

func (p *Payment) GetGoenId() int {
	return p.GoenId
}

func (p *Payment) SetAmountTendered(amountTendered float64) {
	p.AmountTendered = amountTendered
	p.AddBasicFieldChange("amount_tendered")
}
func (p *Payment) GetAmountTendered() float64 {
	return p.AmountTendered
}

func (p *Payment) TurnToCardPayment() (*CardPayment, error) {
	return cardPaymentManager.Get(p.GoenId)
}

type CardPayment struct {
	Payment
	entityManager.FieldChange
	ExpiryDate        time.Time `db:"expiry_date"`
	CardAccountNumber int       `db:"card_account_number"`
}

func (p *CardPayment) GetParentEntity() entityManager.EntityForInheritManager {
	return &p.Payment
}

func (p *CardPayment) SetCardAccountNumber(cardAccountNumber int) {
	p.CardAccountNumber = cardAccountNumber
	p.AddBasicFieldChange("card_account_number")
}
