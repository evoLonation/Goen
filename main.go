package main

import (
	"Cocome/api"
)

func main() {
	config, err := LoadConfig(".")
	if err != nil {
		print("can not load config")
	}

	api.Start(config.ServerAddress)
	//dsn := "root:2002116yy@tcp(127.0.0.1:3306)/Cocome?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//err := db.AutoMigrate(&entityForGen.CardPayment{}, &entityForGen.CashPayment{},
	//	&entityForGen.CashDesk{}, entityForGen.Cashier{},
	//	&entityForGen.Item{}, &entityForGen.OrderProduct{}, &entityForGen.OrderEntry{}, &entityForGen.Supplier{},
	//	entityForGen.ProductCatalog{}, entityForGen.Sale{}, entityForGen.SalesLineItem{}, entityForGen.Store{}, entityForGen.Supplier{})
	//if err != nil {
	//	print(err)
	//	return
	//}
	//operation.CurrentCashDesk = &entity.CashDesk{
	//	IsOpened: true,
	//	Id:       1,
	//}

	//operation.MakeNewSale()
	//operation.EnterItem(1234, 3)

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
//	newItem := entityForGen.Item{
//		Barcode:     1234,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entityForGen.Item{
//		Barcode:     1235,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entityForGen.Item{
//		Barcode:     1236,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entityForGen.Item{
//		Barcode:     1237,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entityForGen.Item{
//		Barcode:     1238,
//		Name:        "生菜",
//		Price:       1.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//	newItem = entityForGen.Item{
//		Barcode:     1239,
//		Name:        "苹果",
//		Price:       2.5,
//		StockNumber: 1000000,
//	}
//	entity.Db.Create(&newItem)
//}
