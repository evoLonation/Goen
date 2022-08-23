package entity

import (
	"Cocome/entityManager"
	"time"
)

var cardPaymentManager entityManager.ManagerForEntity[CardPayment]
var CardPaymentManager entityManager.InheritManagerForOther[CardPayment]

type CardPayment interface {
	Payment
	SetCardAccountNumber(int)
	SetExpiryDate(time.Time)
	GetCardAccountNumber() int
	GetExpiryDate() time.Time
}

type CardPaymentEntity struct {
	PaymentEntity
	entityManager.FieldChange
	ExpiryDate        time.Time `db:"expiry_date"`
	CardAccountNumber int       `db:"card_account_number"`
}

func (p *CardPaymentEntity) GetParentEntity() entityManager.EntityForInheritManager {
	return &p.PaymentEntity
}

func (p *CardPaymentEntity) SetExpiryDate(t time.Time) {
	p.ExpiryDate = t
	p.AddBasicFieldChange("expiry_date")
}

func (p *CardPaymentEntity) SetCardAccountNumber(cardAccountNumber int) {
	p.CardAccountNumber = cardAccountNumber
	p.AddBasicFieldChange("card_account_number")
}

func (p *CardPaymentEntity) GetCardAccountNumber() int {
	return p.CardAccountNumber
}

func (p *CardPaymentEntity) GetExpiryDate() time.Time {
	return p.ExpiryDate
}
