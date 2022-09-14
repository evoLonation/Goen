package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageItemCRUDServiceInstance ManageItemCRUDService

type ManageItemCRUDService struct {
}

func (p *ManageItemCRUDService) createItem(barcode int, name string, price float64, stocknumber int, orderprice float64) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var item entity.Item = entity.ItemRepo.GetFromAllInstanceBy("barcode", barcode)
	// precondition
	if !((item == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var ite entity.Item
	ite = entity.ItemRepo.New()
	ite.SetBarcode(barcode)
	ite.SetName(name)
	ite.SetPrice(price)
	ite.SetStockNumber(stocknumber)
	ite.SetOrderPrice(orderprice)
	entity.ItemRepo.AddInAllInstance(ite)
	ret.Value = true

	return
}
func (p *ManageItemCRUDService) queryItem(barcode int) (ret OperationResult[entity.Item]) {
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
	ret.Value = item

	return
}
func (p *ManageItemCRUDService) modifyItem(barcode int, name string, price float64, stocknumber int, orderprice float64) (ret OperationResult[bool]) {
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
	item.SetBarcode(barcode)
	item.SetName(name)
	item.SetPrice(price)
	item.SetStockNumber(stocknumber)
	item.SetOrderPrice(orderprice)
	ret.Value = true

	return
}
func (p *ManageItemCRUDService) deleteItem(barcode int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var item entity.Item = entity.ItemRepo.GetFromAllInstanceBy("barcode", barcode)
	// precondition
	if !((item == nil) == false && entity.ItemRepo.IsInAllInstance(item)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.ItemRepo.RemoveFromAllInstance(item)
	ret.Value = true

	return
}
