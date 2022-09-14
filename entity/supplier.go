package entity

import (
	"Cocome/entityRepo"
)

var supplierRepo entityRepo.RepoForEntity[Supplier]
var SupplierRepo entityRepo.RepoForOther[Supplier]

type Supplier interface {
	GetId() int
	GetName() string
	SetId(id int)
	SetName(name string)
}

type SupplierEntity struct {
	entityRepo.Entity

	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (p *SupplierEntity) GetId() int {
	return p.Id
}
func (p *SupplierEntity) GetName() string {
	return p.Name
}
func (p *SupplierEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *SupplierEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
