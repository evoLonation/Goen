package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageStoreCRUDServiceInstance ManageStoreCRUDService

type ManageStoreCRUDService struct {
}

func (p *ManageStoreCRUDService) createStore(id int, name string, address string, isopened bool) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var store entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((store == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var sto entity.Store
	sto = entity.StoreRepo.New()
	sto.SetId(id)
	sto.SetName(name)
	sto.SetAddress(address)
	sto.SetIsOpened(isopened)
	entity.StoreRepo.AddInAllInstance(sto)
	ret.Value = true

	return
}
func (p *ManageStoreCRUDService) queryStore(id int) (ret OperationResult[entity.Store]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var store entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((store == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = store

	return
}
func (p *ManageStoreCRUDService) modifyStore(id int, name string, address string, isopened bool) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var store entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((store == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	store.SetId(id)
	store.SetName(name)
	store.SetAddress(address)
	store.SetIsOpened(isopened)
	ret.Value = true

	return
}
func (p *ManageStoreCRUDService) deleteStore(id int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var store entity.Store = entity.StoreRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((store == nil) == false && entity.StoreRepo.IsInAllInstance(store)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.StoreRepo.RemoveFromAllInstance(store)
	ret.Value = true

	return
}
