package entity

import (
	"Cocome/entityManager"
	"log"
)

const (
	CardPaymentInheritType entityManager.GoenInheritType = iota + 1
)

func init() {

	tmpItemManager, err := entityManager.NewManager[ItemEntity, Item]("item")
	if err != nil {
		log.Fatal(err)
	}
	itemManager = tmpItemManager
	ItemManager = tmpItemManager

	tmpPaymentManager, err := entityManager.NewManager[PaymentEntity, Payment]("payment")
	if err != nil {
		log.Fatal(err)
	}
	paymentManager = tmpPaymentManager
	PaymentManager = tmpPaymentManager

	tmpCardPaymentManager, err := entityManager.NewInheritManager[CardPaymentEntity, CardPayment]("card_payment", tmpPaymentManager, CardPaymentInheritType)
	if err != nil {
		log.Fatal(err)
	}
	cardPaymentManager = tmpCardPaymentManager
	CardPaymentManager = tmpCardPaymentManager

}
