package entity

import (
	"Cocome/entityManager"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInherit(t *testing.T) {
	//cardPayment, err := CardPaymentManager.Get(1)
	//payment, err := PaymentManager.Get(1)
	//log.Fatal(err)
	//print(cardPayment)
	//p := CardPaymentManager.New()
	//
	//p.GoenId = 1
	//cp := CardPaymentManager.New()
	//cp.SetAmountTendered(0.111222)
	//cp.SetCardAccountNumber(123124)
	//CardPaymentManager.AddInAllInstance(cp)
	//err := entityManager.Saver.Save()
	//require.NoError(t, err)

	//cp2, err := CardPaymentManager.GetFromAllInstanceBy("amount_tendered", 0.111222)
	cp2, err := CardPaymentManager.GetFromAllInstanceBy("card_account_number", 123124)
	require.NoError(t, err)
	cp2.SetAmountTendered(0.333444)
	cp2.SetCardAccountNumber(3215125)
	err = entityManager.Saver.Save()
	require.NoError(t, err)
	//CardPaymentManager.GetFromAllInstanceBy("expiry_date", 123)
	//CardPaymentManager.GetFromAllInstanceBy("amount_tendered", 123)
	//CardPaymentManager.RemoveFromAllInstance(p)
	//CardPaymentManager.FindFromAllInstanceBy("expiry_date", 123)
	//CardPaymentManager.FindFromAllInstanceBy("amount_tendered", 123)

}
func TestInherit2(t *testing.T) {
	p, err := PaymentManager.GetFromAllInstanceBy("goen_id", 3)
	prr, err := PaymentManager.FindFromAllInstanceBy("goen_id", 3)
	require.NoError(t, err)
	require.NotNil(t, prr)
	var cp *CardPayment
	if p.GetRealType() == CardPaymentInheritType {
		var err error
		cp, err = p.TurnToCardPayment()
		require.NoError(t, err)
	}
	require.NotNil(t, cp)
}

func TestInherit3(t *testing.T) {
	item := ItemManager.New()
	cp := CardPaymentManager.New()
	cp.SetCardAccountNumber(1838940019)
	item.SetPrice(123.0123)
	item.SetBarcode(114514)
	item.SetBelongedPayment(cp)

	err := entityManager.Saver.Save()
	require.NoError(t, err)

	p := item.GetBelongedPayment()
	require.NoError(t, err)
	require.Equal(t, p.GetRealType(), CardPaymentInheritType)
	cp2, err := p.TurnToCardPayment()
	require.NoError(t, err)
	require.NotNil(t, cp2)
}
