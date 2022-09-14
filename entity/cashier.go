package entity

import (
	"Cocome/entityRepo"
)

var cashierRepo entityRepo.RepoForEntity[Cashier]
var CashierRepo entityRepo.RepoForOther[Cashier]

type Cashier interface {
	GetId() int
	GetName() string
	GetWorkedStore() Store
	SetId(id int)
	SetName(name string)
	SetWorkedStore(store Store)
}

type CashierEntity struct {
	entityRepo.Entity

	Id                int    `db:"id"`
	Name              string `db:"name"`
	WorkedStoreGoenId *int   `db:"worked_store_goen_id"`
}

func (p *CashierEntity) GetId() int {
	return p.Id
}
func (p *CashierEntity) GetName() string {
	return p.Name
}
func (p *CashierEntity) GetWorkedStore() Store {
	if p.WorkedStoreGoenId == nil {
		return nil
	} else {
		ret, err := storeRepo.Get(*p.WorkedStoreGoenId)
		if err != nil {
			panic(err)
		}
		return ret
	}
}
func (p *CashierEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *CashierEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *CashierEntity) SetWorkedStore(store Store) {
	id := storeRepo.GetGoenId(store)
	p.WorkedStoreGoenId = &id
	p.AddAssFieldChange("worked_store_goen_id")
}
