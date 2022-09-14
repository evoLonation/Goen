package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var CurrentCashDesk entity.CashDesk
var CurrentStore entity.Store

func openCashDesk(cashDeskID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cd entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", cashDeskID)
	// precondition
	if !((cd == nil) == false && cd.GetIsOpened() == false && (CurrentStore == nil) == false && CurrentStore.GetIsOpened() == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	CurrentCashDesk = cd
	cd.SetIsOpened(true)
	ret.Value = true

	return
}
func closeCashDesk(cashDeskID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cd entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", cashDeskID)
	// precondition
	if !((cd == nil) == false && cd.GetIsOpened() == true && (CurrentStore == nil) == false && CurrentStore.GetIsOpened() == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	CurrentCashDesk = cd
	cd.SetIsOpened(false)
	ret.Value = true

	return
}
func openStore(storeID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var sto entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", storeID)
	// precondition
	if !((sto == nil) == false && sto.GetIsOpened() == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	CurrentStore = sto
	sto.SetIsOpened(true)
	ret.Value = true

	return
}
func closeStore(storeID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var sto entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", storeID)
	// precondition
	if !((sto == nil) == false && sto.GetIsOpened() == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	sto.SetIsOpened(false)
	ret.Value = true

	return
}
func changePrice(barcode int, newPrice float64) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var item entity.Item = entity.ItemRepo.GetFromAllInstanceBy("barcode", barcode)
	// precondition
	if !((item == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	item.SetPrice(newPrice)
	ret.Value = true

	return
}
func receiveOrderedProduct(orderID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var op entity.OrderProduct = entity.OrderProductRepo.GetFromAllInstanceBy("id", orderID)
	// precondition
	if !((op == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	op.SetOrderStatus(entity.OrderStatusRECEIVED)
	ret.Value = true

	return
}
func listSuppliers() (ret OperationResult[[]entity.Supplier]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !(true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = entity.SupplierRepo.GetAll()

	return
}
func showStockReports() (ret OperationResult[[]entity.Item]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !(true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = entity.ItemRepo.GetAll()

	return
}
