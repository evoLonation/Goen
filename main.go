package main

import (
	"Cocome/operation"
	"fmt"
	"gorm.io/gorm/utils"
)

type Store struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	IsOpened bool   `db:"is_opened"`
}

func printStore(store *Store) {
	if store == nil {
		print("store equals to nil! \n")
		return
	}
	fmt.Printf("store.Id : %d\n", store.Id)
	fmt.Printf("store.Name : %s\n", store.Name)
	fmt.Printf("store.Address : %s\n", store.Address)
	fmt.Printf("store.IsOpened: " + utils.ToString(store.IsOpened) + "\n")

}
func main() {
	//如果没有找到...
	var store Store
	store = Store{Id: 123, Address: "giaoge"}
	operation.Db.Get(&store, "select * from store where id = 0")
	printStore(&store)
	//不会改变store的值

	var storep *Store = &Store{}
	operation.Db.Get(storep, "select * from store where id = 0")
	printStore(storep)
	// 会返回nil

	// 如果找到了...
	operation.Db.Get(&store, "select * from store where id = 1")
	printStore(&store)
	//非空的值会修正，为空的值不会
	operation.Db.Get(storep, "select * from store where id = 1")
	printStore(storep)
	// 还是会得到nil

	str := "prte"
	store2 := Store2{Address: &str}
	operation.Db.Get(&store2, "select * from store where id = 1")
	// 会返回正确的值！

	store2.IsOpened = new(bool)
	*store2.IsOpened = true
	*store2.Name = "right"

	operation.Db.NamedExec("update store set is_opened = :is_opened, name = :name", store2)

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
