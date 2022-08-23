package entity

import (
	"Cocome/entityManager"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
	var cp CardPayment
	if PaymentManager.GetRealType(p) == CardPaymentInheritType {
		var err error
		cp, err = CardPaymentManager.CastFrom(p)
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
	require.Equal(t, PaymentManager.GetRealType(p), CardPaymentInheritType)
	cp2, err := CardPaymentManager.CastFrom(p)
	require.NoError(t, err)
	require.NotNil(t, cp2)
}

func TestPayment(t *testing.T) {
	p := PaymentManager.New()
	cp := CardPaymentManager.New()
	cp.SetCardAccountNumber(1838940019)
	p.SetAmountTendered(123)
	cp.SetAmountTendered(456)
	cp.SetExpiryDate(time.Now())
	PaymentManager.AddInAllInstance(p)
	CardPaymentManager.AddInAllInstance(cp)

	err := entityManager.Saver.Save()
	require.NoError(t, err)

	p2, err := PaymentManager.GetFromAllInstanceBy("amount_tendered", 123)
	require.NoError(t, err)
	require.NotNil(t, p2)
	cp2, err := CardPaymentManager.GetFromAllInstanceBy("amount_tendered", 456)
	require.NoError(t, err)
	require.NotNil(t, cp2)
	p3, err := PaymentManager.GetFromAllInstanceBy("amount_tendered", 456)
	require.NoError(t, err)
	require.NotNil(t, p3)

	require.Equal(t, PaymentManager.GetRealType(p3), CardPaymentInheritType)
	cp3, err := CardPaymentManager.CastFrom(p3)
	require.NoError(t, err)
	require.NotNil(t, cp3)

}
