package operation

//
//import (
//	"Cocome/entity"
//	"errors"
//	"time"
//)
//
//var CurrentSaleLine *entity.SalesLineItem
//var CurrentSale *entity.Sale
//var CurrentPaymentMethod *entity.PaymentMethod
//
//func MakeNewSale() (bool, error) {
//
//	// precondition
//	if !(CurrentCashDesk != nil && CurrentCashDesk.IsOpened == true &&
//		(CurrentSale == nil || (CurrentSale != nil && CurrentSale.IsComplete == true))) {
//		return false, errors.New("pre condition dissatisfy")
//	}
//
//	// post condition
//	var s entity.Sale
//	s.BelongedCashDesk = CurrentCashDesk
//	CurrentCashDesk.ContainedSales = append(CurrentCashDesk.ContainedSales, &s)
//	s.IsComplete = false
//	s.IsReadytoPay = false
//	//todo: insert s to Sales
//	CurrentSale = &s
//	return true, nil
//}
//
//func EnterItem(barcode int, quantity int) (bool, error) {
//	// definition
//	var item *entity.Item
//	//todo: get item
//
//	// precondition
//	if !(CurrentSale != nil && CurrentSale.IsComplete == false &&
//		item != nil && item.StockNumber > 0) {
//		return false, errors.New("pre condition dissatisfy")
//	}
//	// post condition
//	var sli entity.SalesLineItem
//	CurrentSaleLine = &sli
//	sli.BelongedSale = CurrentSale
//	CurrentSale.ContainedSalesLine = append(CurrentSale.ContainedSalesLine, &sli)
//	sli.Quantity = quantity
//	sli.BelongedItem = item
//	item.StockNumber -= quantity
//	sli.Subamount = item.Price * float64(quantity)
//	//todo: insert sli to SaleLines
//	return true, nil
//}
//
//func EndSale() (float64, error) {
//	// definition
//	var sls []entity.SalesLineItem
//	var sub []float64
//	//todo: get sls
//	for i := 0; i <= len(sls); i++ {
//		sub = append(sub, sls[i].Subamount)
//	}
//
//	// precondition
//	if !(CurrentSale != nil && CurrentSale.IsComplete == false &&
//		CurrentSale.IsReadytoPay == false) {
//		return 0, errors.New("pre condition dissatisfy")
//	}
//
//	// post condition
//	CurrentSale.Amount = Sum(sub)
//	CurrentSale.IsReadytoPay = true
//	return CurrentSale.Amount, nil
//}
//
//type Addable interface {
//	int | float32 | float64 | string
//}
//
//func Sum[T Addable](arr []T) T {
//	var a T
//	for i := 0; i <= len(arr); i++ {
//		a = arr[i] + a
//	}
//	return a
//}
//
//func MakeCashPayment(amount float64) (bool, error) {
//	// precondition
//	if !(CurrentSale != nil &&
//		CurrentSale.IsComplete == false &&
//		CurrentSale.IsReadytoPay == true &&
//		amount >= CurrentSale.Amount) {
//		return false, errors.New("pre condition dissatisfy")
//	}
//
//	// post condition
//	var cp entity.CashPayment
//	cp.AmountTendered = amount
//	cp.BelongedSale = CurrentSale
//	// golang can not Inheritance !!!
//	CurrentSale.AssoicatedCashPayment = &cp
//	CurrentSale.BelongedStore = CurrentStore
//	CurrentStore.Sales = append(CurrentStore.Sales, CurrentSale)
//	CurrentSale.IsComplete = true
//	CurrentSale.Time = time.Now()
//	cp.Balance = amount - CurrentSale.Amount
//	// todo: insert cp
//	return true, nil
//}
//
//func MakeCardPayment(cardAccountNumber string, expiryDate time.Time, fee float64) (bool, error) {
//	// precondition
//	//todo: implement thirdPartyCardPaymentService(cardAccountNumber, expiryDate, fee)
//	if !(CurrentSale != nil &&
//		CurrentSale.IsComplete == false &&
//		CurrentSale.IsReadytoPay == true) {
//		return false, errors.New("pre condition dissatisfy")
//	}
//
//	// post condition
//	var cdp entity.CardPayment
//	cdp.AmountTendered = fee
//	cdp.BelongedSale = CurrentSale
//	// golang can not Inheritance !!!
//	CurrentSale.AssoicatedCardPayment = &cdp
//	cdp.CardAccountNumber = cardAccountNumber
//	cdp.ExpiryDate = expiryDate
//	// todo: insert cdp
//	CurrentSale.BelongedStore = CurrentStore
//	CurrentStore.Sales = append(CurrentStore.Sales, CurrentSale)
//	CurrentSale.IsComplete = true
//	CurrentSale.Time = time.Now()
//	return true, nil
//}
