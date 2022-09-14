package entity

import (
	"Cocome/entityRepo"
)

var cashPaymentRepo entityRepo.RepoForEntity[CashPayment]
var CashPaymentRepo entityRepo.InheritRepoForOther[CashPayment]

type CashPayment interface {
	Payment
	GetBalance() float64
	SetBalance(balance float64)
}

type CashPaymentEntity struct {
	PaymentEntity
	entityRepo.FieldChange

	Balance float64 `db:"balance"`
}

func (p *CashPaymentEntity) GetParentEntity() entityRepo.EntityForInheritRepo {
	return &p.PaymentEntity
}

func (p *CashPaymentEntity) GetBalance() float64 {
	return p.Balance
}
func (p *CashPaymentEntity) SetBalance(balance float64) {
	p.Balance = balance
	p.AddBasicFieldChange("balance")
}
