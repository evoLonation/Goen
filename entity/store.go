package entity

import (
	"Cocome/entityRepo"
)

var storeRepo entityRepo.RepoForEntity[Store]
var StoreRepo entityRepo.RepoForOther[Store]

type Store interface {
	GetId() int
	GetName() string
	GetAddress() string
	GetIsOpened() bool
	GetAssociationCashdeskes() []CashDesk
	GetProductcatalogs() []ProductCatalog
	GetItems() []Item
	GetCashiers() []Cashier
	GetSales() []Sale
	SetId(id int)
	SetName(name string)
	SetAddress(address string)
	SetIsOpened(isOpened bool)
	AddAssociationCashdeskes(cashDesk CashDesk)
	RemoveAssociationCashdeskes(cashDesk CashDesk)
	AddProductcatalogs(productCatalog ProductCatalog)
	RemoveProductcatalogs(productCatalog ProductCatalog)
	AddItems(item Item)
	RemoveItems(item Item)
	AddCashiers(cashier Cashier)
	RemoveCashiers(cashier Cashier)
	AddSales(sale Sale)
	RemoveSales(sale Sale)
}

type StoreEntity struct {
	entityRepo.Entity

	Id       int    `db:"id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	IsOpened bool   `db:"is_opened"`
}

func (p *StoreEntity) GetId() int {
	return p.Id
}
func (p *StoreEntity) GetName() string {
	return p.Name
}
func (p *StoreEntity) GetAddress() string {
	return p.Address
}
func (p *StoreEntity) GetIsOpened() bool {
	return p.IsOpened
}
func (p *StoreEntity) GetAssociationCashdeskes() []CashDesk {
	ret, _ := cashDeskRepo.FindFromMultiAssTable("store_association_cashdeskes", p.GoenId)
	return ret
}
func (p *StoreEntity) GetProductcatalogs() []ProductCatalog {
	ret, _ := productCatalogRepo.FindFromMultiAssTable("store_productcatalogs", p.GoenId)
	return ret
}
func (p *StoreEntity) GetItems() []Item {
	ret, _ := itemRepo.FindFromMultiAssTable("store_items", p.GoenId)
	return ret
}
func (p *StoreEntity) GetCashiers() []Cashier {
	ret, _ := cashierRepo.FindFromMultiAssTable("store_cashiers", p.GoenId)
	return ret
}
func (p *StoreEntity) GetSales() []Sale {
	ret, _ := saleRepo.FindFromMultiAssTable("store_sales", p.GoenId)
	return ret
}
func (p *StoreEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *StoreEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *StoreEntity) SetAddress(address string) {
	p.Address = address
	p.AddBasicFieldChange("address")
}
func (p *StoreEntity) SetIsOpened(isOpened bool) {
	p.IsOpened = isOpened
	p.AddBasicFieldChange("is_opened")
}
func (p *StoreEntity) AddAssociationCashdeskes(cashDesk CashDesk) {
	p.AddMultiAssChange(entityRepo.Include, "store_association_cashdeskes", cashDeskRepo.GetGoenId(cashDesk))
}
func (p *StoreEntity) RemoveAssociationCashdeskes(cashDesk CashDesk) {
	p.AddMultiAssChange(entityRepo.Exclude, "store_association_cashdeskes", cashDeskRepo.GetGoenId(cashDesk))
}
func (p *StoreEntity) AddProductcatalogs(productCatalog ProductCatalog) {
	p.AddMultiAssChange(entityRepo.Include, "store_productcatalogs", productCatalogRepo.GetGoenId(productCatalog))
}
func (p *StoreEntity) RemoveProductcatalogs(productCatalog ProductCatalog) {
	p.AddMultiAssChange(entityRepo.Exclude, "store_productcatalogs", productCatalogRepo.GetGoenId(productCatalog))
}
func (p *StoreEntity) AddItems(item Item) {
	p.AddMultiAssChange(entityRepo.Include, "store_items", itemRepo.GetGoenId(item))
}
func (p *StoreEntity) RemoveItems(item Item) {
	p.AddMultiAssChange(entityRepo.Exclude, "store_items", itemRepo.GetGoenId(item))
}
func (p *StoreEntity) AddCashiers(cashier Cashier) {
	p.AddMultiAssChange(entityRepo.Include, "store_cashiers", cashierRepo.GetGoenId(cashier))
}
func (p *StoreEntity) RemoveCashiers(cashier Cashier) {
	p.AddMultiAssChange(entityRepo.Exclude, "store_cashiers", cashierRepo.GetGoenId(cashier))
}
func (p *StoreEntity) AddSales(sale Sale) {
	p.AddMultiAssChange(entityRepo.Include, "store_sales", saleRepo.GetGoenId(sale))
}
func (p *StoreEntity) RemoveSales(sale Sale) {
	p.AddMultiAssChange(entityRepo.Exclude, "store_sales", saleRepo.GetGoenId(sale))
}
