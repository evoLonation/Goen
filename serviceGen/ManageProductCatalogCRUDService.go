package serviceGen

import (
	"Cocome/entity"
	"Cocome/entityRepo"
)

var ManageProductCatalogCRUDServiceInstance ManageProductCatalogCRUDService

type ManageProductCatalogCRUDService struct {
}

func (p *ManageProductCatalogCRUDService) createProductCatalog(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var productcatalog entity.ProductCatalog = entity.ProductCatalogRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((productcatalog == nil) == true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	var pro entity.ProductCatalog
	pro = entity.ProductCatalogRepo.New()
	pro.SetId(id)
	pro.SetName(name)
	entity.ProductCatalogRepo.AddInAllInstance(pro)
	ret.Value = true

	return
}
func (p *ManageProductCatalogCRUDService) queryProductCatalog(id int) (ret OperationResult[entity.ProductCatalog]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var productcatalog entity.ProductCatalog = entity.ProductCatalogRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((productcatalog == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = productcatalog

	return
}
func (p *ManageProductCatalogCRUDService) modifyProductCatalog(id int, name string) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var productcatalog entity.ProductCatalog = entity.ProductCatalogRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((productcatalog == nil) == false) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	productcatalog.SetId(id)
	productcatalog.SetName(name)
	ret.Value = true

	return
}
func (p *ManageProductCatalogCRUDService) deleteProductCatalog(id int) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	//definition
	var productcatalog entity.ProductCatalog = entity.ProductCatalogRepo.GetFromAllInstanceBy("id", id)
	// precondition
	if !((productcatalog == nil) == false && entity.ProductCatalogRepo.IsInAllInstance(productcatalog)) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	entity.ProductCatalogRepo.RemoveFromAllInstance(productcatalog)
	ret.Value = true

	return
}
