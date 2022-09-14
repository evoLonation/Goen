package entity

import (
	"Cocome/entityRepo"
)

var orderEntryRepo entityRepo.RepoForEntity[OrderEntry]
var OrderEntryRepo entityRepo.RepoForOther[OrderEntry]

type OrderEntry interface {
	GetQuantity() int
	GetSubAmount() float64
	GetItem() Item
	SetQuantity(quantity int)
	SetSubAmount(subAmount float64)
	SetItem(item Item)
}

type OrderEntryEntity struct {
	entityRepo.Entity

	Quantity   int     `db:"quantity"`
	SubAmount  float64 `db:"sub_amount"`
	ItemGoenId *int    `db:"item_goen_id"`
}

func (p *OrderEntryEntity) GetQuantity() int {
	return p.Quantity
}
func (p *OrderEntryEntity) GetSubAmount() float64 {
	return p.SubAmount
}
func (p *OrderEntryEntity) GetItem() Item {
	if p.ItemGoenId == nil {
		return nil
	} else {
		ret, err := itemRepo.Get(*p.ItemGoenId)
		if err != nil {
			panic(err)
		}
		return ret
	}
}
func (p *OrderEntryEntity) SetQuantity(quantity int) {
	p.Quantity = quantity
	p.AddBasicFieldChange("quantity")
}
func (p *OrderEntryEntity) SetSubAmount(subAmount float64) {
	p.SubAmount = subAmount
	p.AddBasicFieldChange("sub_amount")
}
func (p *OrderEntryEntity) SetItem(item Item) {
	id := itemRepo.GetGoenId(item)
	p.ItemGoenId = &id
	p.AddAssFieldChange("item_goen_id")
}
