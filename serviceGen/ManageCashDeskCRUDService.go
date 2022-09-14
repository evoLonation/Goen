package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageCashDeskCRUDServiceInstance ManageCashDeskCRUDService

type ManageCashDeskCRUDService struct {
}

func (p *ManageCashDeskCRUDService) createCashDesk(id int, name string, isopened bool) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashdesk entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashdesk == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var cas entity.CashDesk
	cas = entity.CashDeskRepo.New()
	cas.SetId(id)
	cas.SetName(name)
	cas.SetIsOpened(isopened)
	entity.CashDeskRepo.AddInAllInstance(cas)
	ret.Value = true

	return
}
func (p *ManageCashDeskCRUDService) queryCashDesk(id int) (ret OperationResult[entity.CashDesk]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashdesk entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashdesk == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = cashdesk

	return
}
func (p *ManageCashDeskCRUDService) modifyCashDesk(id int, name string, isopened bool) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashdesk entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashdesk == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	cashdesk.SetId(id)
	cashdesk.SetName(name)
	cashdesk.SetIsOpened(isopened)
	ret.Value = true

	return
}
func (p *ManageCashDeskCRUDService) deleteCashDesk(id int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var cashdesk entity.CashDesk = entity.CashDeskRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((cashdesk == nil) == false && entity.CashDeskRepo.IsInAllInstance(cashdesk)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.CashDeskRepo.RemoveFromAllInstance(cashdesk)
	ret.Value = true

	return
}
