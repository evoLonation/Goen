package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
	"time"
)

var CoCoMEOrderProductsInstance CoCoMEOrderProducts

type CoCoMEOrderProducts struct {
	CurrentOrderProduct entity.OrderProduct
}

func (p *CoCoMEOrderProducts) makeNewOrder(orderid int) (ret OperationResult[bool]) {
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
	var op entity.OrderProduct
	op = entity.OrderProductRepo.New()
	op.SetOrderStatus(entity.OrderStatusNEW)
	op.SetId(orderid)
	op.SetTime(time.Now())
	entity.OrderProductRepo.AddInAllInstance(op)
	p.CurrentOrderProduct = op
	ret.Value = true

	return
}
func (p *CoCoMEOrderProducts) listAllOutOfStoreProducts() (ret OperationResult[[]entity.Item]) {
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
	ret.Value = entity.ItemRepo.FindFromAllInstanceBy("stock_number", 0)

	return
}
func (p *CoCoMEOrderProducts) orderItem(barcode int, quantity int) (ret OperationResult[bool]) {
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
	var order entity.OrderEntry
	order = entity.OrderEntryRepo.New()
	order.SetQuantity(quantity)
	order.SetSubAmount(item.GetOrderPrice() * float64(quantity))
	order.SetItem(item)
	entity.OrderEntryRepo.AddInAllInstance(order)
	p.CurrentOrderProduct.AddContainedEntries(order)
	ret.Value = true

	return
}
func (p *CoCoMEOrderProducts) chooseSupplier(supplierID int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var sup entity.Supplier = entity.SupplierRepo.GetFromAllInstanceBy("id", supplierID)
	// precondition
	if !((sup == nil) == false && (p.CurrentOrderProduct == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	p.CurrentOrderProduct.SetSupplier(sup)
	ret.Value = true

	return
}
func (p *CoCoMEOrderProducts) placeOrder() (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !((p.CurrentOrderProduct == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	p.CurrentOrderProduct.SetOrderStatus(entity.OrderStatusREQUESTED)
	ret.Value = true

	return
}
