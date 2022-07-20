package entityForGen

//
//import (
//	"time"
//)
//
//type Item struct {
//	Barcode     int     `gorm:"primaryKey;autoIncrement"`
//	Name        string  `gorm:"not null"`
//	Price       float64 `gorm:"not null"`
//	StockNumber int     `gorm:"not null"`
//}
//
//type Cashier struct {
//	Id   int
//	Name string
//}
//
//// Cart
////一个购物车拥有很多Item,Has Many
////一个购物车拥有一个Payment, Has One
////Time是创建时间
////isComplete代表是否支付完毕
//type Cart struct {
//	Id         int
//	Time       time.Time `gorm:"autoCreateTime:int"`
//	IsComplete bool      `gorm:"default:false"`
//	Amount     float64   `gorm:"default:0"`
//	Items      []ItemInCart
//	Payment    *Payment
//}
//
//// ItemInCart 属于一个item
////为什么不是Item拥有很多ItemInCart？因为ItemInShopping需要知道它的Item，但是Item不需要频繁查找它的所有ItemInShopping
//// 创建该对象时必须关联一个item，并且自动设置数量为1
//type ItemInCart struct {
//	Id        int
//	Quantity  int `gorm:"default:1"`
//	SubAmount float64
//	CartId    int `gorm:"not null"`
//	Item      *Item
//	ItemId    int `gorm:"not null"`
//}
//
////// AfterCreate 如果出错钩子会默认回滚
////func (i *ItemInCart) AfterCreate(tx *gorm.DB) (err error) {
////	if i.Item.StockNumber > 0 {
////		return errors.New("该商品数量不足1，无法加入购物车！")
////	}
////	i.Item.StockNumber--
////	i.SubAmount = i.Item.Price
////	//使用FullSaveAssociations保存所有关联的改动
////	tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&i)
////	return
////}
////func (i *ItemInCart) BeforeCreate(tx *gorm.DB) (err error) {
////	if tx.Model(Item{}).First(i.ItemId).Error != nil {
////		return errors.New("无法找到该条形码的商品！")
////	}
////	return
////}
//
//// Payment 属于一个信用卡
//type Payment struct {
//	Id             int
//	AmountTendered float64
//	Balance        float64
//	CartId         string
//	CreditCard     *CreditCard
//	CreditCardId   string
//}
//
//type CreditCard struct {
//	Id      string  `gorm:"primaryKey;size:5"`
//	Balance float64 `gorm:"not null"`
//}
