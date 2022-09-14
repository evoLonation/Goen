package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
	"Cocome/util"
	"time"
)

var ProcessSaleServiceInstance ProcessSaleService

type ProcessSaleService struct {
	CurrentSaleLine      entity.SalesLineItem
	CurrentSale          entity.Sale
	CurrentPaymentMethod entity.PaymentMethod
}

func (p *ProcessSaleService) makeNewSale() (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !((CurrentCashDesk == nil) == false && CurrentCashDesk.GetIsOpened() == true && ((p.CurrentSale == nil) == true || ((p.CurrentSale == nil) == false && p.CurrentSale.GetIsComplete() == true))) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var s entity.Sale
	s = entity.SaleRepo.New()
	s.SetBelongedCashDesk(CurrentCashDesk)
	CurrentCashDesk.AddContainedSales(s)
	s.SetIsComplete(false)
	s.SetIsReadytoPay(false)
	entity.SaleRepo.AddInAllInstance(s)
	p.CurrentSale = s
	ret.Value = true

	return
}
func (p *ProcessSaleService) enterItem(barcode int, quantity int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var item entity.Item = entity.ItemRepo.GetFromAllInstanceBy("barcode", barcode)
	// precondition
	if !((p.CurrentSale == nil) == false && p.CurrentSale.GetIsComplete() == false && (item == nil) == false && item.GetStockNumber() > 0) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var sli entity.SalesLineItem
	sli = entity.SalesLineItemRepo.New()
	p.CurrentSaleLine = sli
	sli.SetBelongedSale(p.CurrentSale)
	p.CurrentSale.AddContainedSalesLine(sli)
	sli.SetQuantity(quantity)
	sli.SetBelongedItem(item)
	item.SetStockNumber(item.GetStockNumber() - quantity)
	sli.SetSubamount(item.GetPrice() * float64(quantity))
	entity.SalesLineItemRepo.AddInAllInstance(sli)
	ret.Value = true

	return
}
func (p *ProcessSaleService) endSale() (ret OperationResult[float64]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var sls []entity.SalesLineItem = p.CurrentSale.GetContainedSalesLine()
	var sub []float64 = util.Collect(sls, func(s entity.SalesLineItem) float64 { return s.GetSubamount() })
	// precondition
	if !((p.CurrentSale == nil) == false && p.CurrentSale.GetIsComplete() == false && p.CurrentSale.GetIsReadytoPay() == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	p.CurrentSale.SetAmount(util.Sum(sub))
	p.CurrentSale.SetIsReadytoPay(true)
	ret.Value = p.CurrentSale.GetAmount()

	return
}
func (p *ProcessSaleService) makeCashPayment(amount float64) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !((p.CurrentSale == nil) == false && p.CurrentSale.GetIsComplete() == false && p.CurrentSale.GetIsReadytoPay() == true && amount >= p.CurrentSale.GetAmount()) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var cp entity.CashPayment
	cp = entity.CashPaymentRepo.New()
	cp.SetAmountTendered(amount)
	cp.SetBelongedSale(p.CurrentSale)
	p.CurrentSale.SetAssoicatedPayment(cp)
	p.CurrentSale.SetBelongedstore(CurrentStore)
	CurrentStore.AddSales(p.CurrentSale)
	p.CurrentSale.SetIsComplete(true)
	p.CurrentSale.SetTime(time.Now())
	cp.SetBalance(amount - p.CurrentSale.GetAmount())
	entity.CashPaymentRepo.AddInAllInstance(cp)
	ret.Value = true

	return
}
func (p *ProcessSaleService) makeCardPayment(cardAccountNumber string, expiryDate time.Time, fee float64) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !((p.CurrentSale == nil) == false && p.CurrentSale.GetIsComplete() == false && p.CurrentSale.GetIsReadytoPay() == true && ThirdPartyServicesInstance.thirdPartyCardPaymentService(cardAccountNumber, expiryDate, fee).Value) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var cdp entity.CardPayment
	cdp = entity.CardPaymentRepo.New()
	cdp.SetAmountTendered(fee)
	cdp.SetBelongedSale(p.CurrentSale)
	p.CurrentSale.SetAssoicatedPayment(cdp)
	cdp.SetCardAccountNumber(cardAccountNumber)
	cdp.SetExpiryDate(expiryDate)
	entity.CardPaymentRepo.AddInAllInstance(cdp)
	p.CurrentSale.SetBelongedstore(CurrentStore)
	CurrentStore.AddSales(p.CurrentSale)
	p.CurrentSale.SetIsComplete(true)
	p.CurrentSale.SetTime(time.Now())
	ret.Value = true

	return
}
