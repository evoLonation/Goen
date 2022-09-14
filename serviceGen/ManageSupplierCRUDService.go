package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageSupplierCRUDServiceInstance ManageSupplierCRUDService

type ManageSupplierCRUDService struct {
}

func (p *ManageSupplierCRUDService) createSupplier(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var supplier entity.Supplier = entity.SupplierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((supplier == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var sup entity.Supplier
	sup = entity.SupplierRepo.New()
	sup.SetId(id)
	sup.SetName(name)
	entity.SupplierRepo.AddInAllInstance(sup)
	ret.Value = true

	return
}
func (p *ManageSupplierCRUDService) querySupplier(id int) (ret OperationResult[entity.Supplier]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var supplier entity.Supplier = entity.SupplierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((supplier == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = supplier

	return
}
func (p *ManageSupplierCRUDService) modifySupplier(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var supplier entity.Supplier = entity.SupplierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((supplier == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	supplier.SetId(id)
	supplier.SetName(name)
	ret.Value = true

	return
}
func (p *ManageSupplierCRUDService) deleteSupplier(id int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var supplier entity.Supplier = entity.SupplierRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((supplier == nil) == false && entity.SupplierRepo.IsInAllInstance(supplier)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.SupplierRepo.RemoveFromAllInstance(supplier)
	ret.Value = true

	return
}
