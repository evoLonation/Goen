package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageCashierCRUDServiceInstance ManageCashierCRUDService

type ManageCashierCRUDService struct {
}

func (p *ManageCashierCRUDService) createCashier(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashier entity.Cashier = entity.CashierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashier == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var cas entity.Cashier
	cas = entity.CashierRepo.New()
	cas.SetId(id)
	cas.SetName(name)
	entity.CashierRepo.AddInAllInstance(cas)
	ret.Value = true

	return
}
func (p *ManageCashierCRUDService) queryCashier(id int) (ret OperationResult[entity.Cashier]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashier entity.Cashier = entity.CashierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashier == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = cashier

	return
}
func (p *ManageCashierCRUDService) modifyCashier(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashier entity.Cashier = entity.CashierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashier == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	cashier.SetId(id)
	cashier.SetName(name)
	ret.Value = true

	return
}
func (p *ManageCashierCRUDService) deleteCashier(id int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashier entity.Cashier = entity.CashierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashier == nil) == false && entity.CashierRepo.IsInAllInstance(cashier)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.CashierRepo.RemoveFromAllInstance(cashier)
	ret.Value = true

	return
}
