package operation

import (
	"Cocome/entity"
	"database/sql"
	"errors"
	"time"
)

var CurrentSaleLine *entity.SalesLineItem
var CurrentSale *entity.Sale
var CurrentPaymentMethod *entity.PaymentMethod

func MakeNewSale() (bool, error) {

	//precondition
	if !(CurrentCashDesk != nil && CurrentCashDesk.IsOpened == true &&
		(CurrentSale == nil || (CurrentSale != nil && CurrentSale.IsComplete == true))) {
		return false, errors.New("pre condition dissatisfy")
	}

	// post condition
	var s entity.Sale
	s.BelongedCashDesk = CurrentCashDesk
	CurrentCashDesk.ContainedSales = append(CurrentCashDesk.ContainedSales, &s)
	s.CashDeskId = sql.NullInt32{Int32: int32(CurrentCashDesk.Id), Valid: true}
	s.IsComplete = false
	s.IsReadytoPay = false
	_, err := db.NamedExec(`insert into sales (time, is_complete, amount, is_readyto_pay, cash_desk_id) 
									values (:Time, :IsComplete, :Amount, :IsReadytoPay, :CashDeskId)`, &s)
	if err != nil {
		print(err.Error())
	}
	CurrentSale = &s
	return true, nil
}

func EnterItem(barcode int, quantity int) (bool, error) {
	// definition
	item := &entity.Item{}
	err := db.Get(item, "select * from items where barcode=?", barcode)
	if err != nil {
		print(err.Error())
		return false, err
	}

	// precondition
	if !(CurrentSale != nil && CurrentSale.IsComplete == false &&
		item != nil && item.StockNumber > 0) {
		return false, errors.New("pre condition dissatisfy")
	}

	// post condition
	var sli entity.SalesLineItem
	CurrentSaleLine = &sli
	sli.BelongedSale = CurrentSale
	CurrentSale.ContainedSalesLine = append(CurrentSale.ContainedSalesLine, &sli)
	sli.SaleId = CurrentSale.GoenId

	sli.Quantity = quantity
	sli.BelongedItem = item
	sli.ItemId = item.Barcode
	item.StockNumber -= quantity
	_, err = db.Exec(`update items set stock_number = ? where barcode = ?`, item.StockNumber, item.Barcode)
	if err != nil {
		print(err.Error())
		return false, err
	}
	sli.Subamount = item.Price * float64(quantity)
	_, err = db.NamedExec(`insert into sales_line_items (quantity, subamount, sale_id, item_id) 
									values (:quantity, :subamount, :sale_id, :item_id)`, &sli)
	if err != nil {
		print(err.Error())
		return false, err
	}
	return true, nil
}

func EndSale() (sql.NullFloat64, error) {
	// definition
	var sls []entity.SalesLineItem
	var sub []float64
	//todo: get sls
	for i := 0; i <= len(sls); i++ {
		sub = append(sub, sls[i].Subamount)
	}

	// precondition
	if !(CurrentSale != nil && CurrentSale.IsComplete == false &&
		CurrentSale.IsReadytoPay == false) {
		return sql.NullFloat64{}, errors.New("pre condition dissatisfy")
	}

	// post condition
	CurrentSale.Amount = sql.NullFloat64{Float64: Sum(sub)}
	CurrentSale.IsReadytoPay = true
	return CurrentSale.Amount, nil
}

type Addable interface {
	int | float32 | float64 | string
}

func Sum[T Addable](arr []T) T {
	var a T
	for i := 0; i <= len(arr); i++ {
		a = arr[i] + a
	}
	return a
}

func MakeCashPayment(amount float64) (bool, error) {
	// precondition
	if !(CurrentSale != nil &&
		CurrentSale.IsComplete == false &&
		CurrentSale.IsReadytoPay == true &&
		amount >= CurrentSale.Amount.Float64) {
		return false, errors.New("pre condition dissatisfy")
	}

	// post condition
	var cp entity.CashPayment
	cp.AmountTendered = amount
	cp.BelongedSale = CurrentSale
	// golang can not Inheritance !!!
	CurrentSale.AssoicatedCashPayment = &cp
	CurrentSale.BelongedStore = CurrentStore
	CurrentStore.Sales = append(CurrentStore.Sales, CurrentSale)
	CurrentSale.IsComplete = true
	CurrentSale.Time = sql.NullTime{Time: time.Now()}
	cp.Balance = sql.NullFloat64{Float64: amount - CurrentSale.Amount.Float64}
	// todo: insert cp
	return true, nil
}

func MakeCardPayment(cardAccountNumber string, expiryDate time.Time, fee float64) (bool, error) {
	// precondition
	//todo: implement thirdPartyCardPaymentService(cardAccountNumber, expiryDate, fee)
	if !(CurrentSale != nil &&
		CurrentSale.IsComplete == false &&
		CurrentSale.IsReadytoPay == true) {
		return false, errors.New("pre condition dissatisfy")
	}

	// post condition
	var cdp entity.CardPayment
	cdp.AmountTendered = fee
	cdp.BelongedSale = CurrentSale
	// golang can not Inheritance !!!
	CurrentSale.AssoicatedCardPayment = &cdp
	cdp.CardAccountNumber = cardAccountNumber
	cdp.ExpiryDate = expiryDate
	// todo: insert cdp
	CurrentSale.BelongedStore = CurrentStore
	CurrentStore.Sales = append(CurrentStore.Sales, CurrentSale)
	CurrentSale.IsComplete = true
	//CurrentSale.Time = time.Now()
	return true, nil
}
