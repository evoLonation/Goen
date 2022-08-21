package entity

import "Cocome/entityManager"

var itemManager entityManager.ManagerForEntity[*ItemEntity] = entityManager.NewManager[ItemEntity]("item")
var ItemManager entityManager.ManagerForOther[*ItemEntity] = itemManager.(*entityManager.Manager[ItemEntity, *ItemEntity])

type ItemInterface interface {
	SetName(name string)
	SetBarcode(barcode int)
	SetPrice(price float64)
	SetOrderPrice(price float64)
	SetStockNumber(stockNumber int)
	AddContainedItem(item *ItemEntity)
	SetBelongedItem(item *ItemEntity)
	SetBelongedPayment(payment PaymentInterface)

	GetName() string
	GetBarcode() int
	GetPrice() float64
	GetOrderPrice() float64
	GetStockNumber() int
	GetContainedItem() []*ItemEntity
	GetBelongedItem() *ItemEntity
	GetBelongedPayment() PaymentInterface
}

type ItemEntity struct {
	entityManager.Entity
	Barcode     int     `db:"barcode"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	StockNumber int     `db:"stock_number"`
	OrderPrice  float64 `db:"order_price"`

	BelongedItemGoenId    *int `db:"belonged_item_goen_id"`
	BelongedPaymentGoenId *int `db:"belonged_payment_goen_id"`
}

func (p *ItemEntity) GetName() string {
	return p.Name
}

func (p *ItemEntity) GetBarcode() int {
	return p.Barcode
}

func (p *ItemEntity) GetPrice() float64 {
	return p.Price
}

func (p *ItemEntity) GetOrderPrice() float64 {
	return p.OrderPrice
}

func (p *ItemEntity) GetStockNumber() int {
	return p.StockNumber
}

func (p *ItemEntity) GetContainedItem() []*ItemEntity {
	ret, _ := itemManager.FindFromMultiAssTable("item_contained_item", p.GoenId)
	return ret
}

func (p *ItemEntity) GetBelongedItem() *ItemEntity {
	if p.BelongedItemGoenId == nil {
		return nil
	} else {
		ret, _ := itemManager.Get(*p.BelongedItemGoenId)
		return ret
	}
}

func (p *ItemEntity) GetBelongedPayment() PaymentInterface {
	if p.BelongedPaymentGoenId == nil {
		return nil
	} else {
		ret, _ := paymentManager.Get(*p.BelongedPaymentGoenId)
		return ret
	}
}

func (p *ItemEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *ItemEntity) SetBarcode(barcode int) {
	p.Barcode = barcode
	p.AddBasicFieldChange("barcode")
}
func (p *ItemEntity) SetPrice(price float64) {
	p.Price = price
	p.AddBasicFieldChange("price")
}
func (p *ItemEntity) SetOrderPrice(price float64) {
	p.OrderPrice = price
	p.AddBasicFieldChange("order_price")
}

func (p *ItemEntity) SetStockNumber(stockNumber int) {
	p.StockNumber = stockNumber
	p.AddBasicFieldChange("stock_number")
}

func (p *ItemEntity) AddContainedItem(item *ItemEntity) {
	p.AddMultiAssChange(entityManager.Include, "item_contained_item", item.GoenId)
}
func (p *ItemEntity) SetBelongedItem(item *ItemEntity) {
	p.BelongedItemGoenId = &item.GoenId
	p.AddAssFieldChange("belonged_item_goen_id")
}

func (p *ItemEntity) SetBelongedPayment(payment PaymentInterface) {
	goenId := payment.GetGoenId()
	p.BelongedPaymentGoenId = &goenId
	p.AddAssFieldChange("belonged_payment_goen_id")
}
