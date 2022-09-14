package entity

import (
	"Cocome/entityRepo"
)

var productCatalogRepo entityRepo.RepoForEntity[ProductCatalog]
var ProductCatalogRepo entityRepo.RepoForOther[ProductCatalog]

type ProductCatalog interface {
	GetId() int
	GetName() string
	GetContainedItems() []Item
	SetId(id int)
	SetName(name string)
	AddContainedItems(item Item)
	RemoveContainedItems(item Item)
}

type ProductCatalogEntity struct {
	entityRepo.Entity

	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (p *ProductCatalogEntity) GetId() int {
	return p.Id
}
func (p *ProductCatalogEntity) GetName() string {
	return p.Name
}
func (p *ProductCatalogEntity) GetContainedItems() []Item {
	ret, _ := itemRepo.FindFromMultiAssTable("product_catalog_contained_items", p.GoenId)
	return ret
}
func (p *ProductCatalogEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *ProductCatalogEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *ProductCatalogEntity) AddContainedItems(item Item) {
	p.AddMultiAssChange(entityRepo.Include, "product_catalog_contained_items", itemRepo.GetGoenId(item))
}
func (p *ProductCatalogEntity) RemoveContainedItems(item Item) {
	p.AddMultiAssChange(entityRepo.Exclude, "product_catalog_contained_items", itemRepo.GetGoenId(item))
}
