package main

import (
	"Cocome/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//operation.CurrentSale = nil
	dsn := "root:2002116yy@tcp(127.0.0.1:3306)/Cocome?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err := db.AutoMigrate(&entity.CardPayment{}, &entity.CashPayment{},
		&entity.CashDesk{}, entity.Cashier{},
		&entity.Item{}, &entity.OrderProduct{}, &entity.OrderEntry{}, &entity.Supplier{},
		entity.ProductCatalog{}, entity.Sale{}, entity.SalesLineItem{}, entity.Store{}, entity.Supplier{})
	if err != nil {
		print(err)
		return
	}
	//createItems()
	//operation.MakeNewShoppingCart()
	//controller.Start()
	//entity.Db.Create(&newItem)
	//operation.AddItem(1234)
	//operation.AddItem(1235)
	//entity.Db.Create(&newItem)
	//
	/////////////ProcessSale部分
	//operation.MakeNewSale()
	//operation.EnterItem(1234, 10)
	//operation.EnterItem(1235, 20)
	//operation.EndSale()
	//operation.MakeCashPayment(100)
}

//func createItems() {
//	newItem := entity.Item{
//		Barcode:     1234,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entity.Item{
//		Barcode:     1235,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entity.Item{
//		Barcode:     1236,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entity.Item{
//		Barcode:     1237,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entity.Item{
//		Barcode:     1238,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entity.Item{
//		Barcode:     1239,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//}
