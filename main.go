package main

import (
	"Cocome/entityManager"
	"log"
)

func main() {
	//item := entityManager.ItemManager.New()
	//item.SetBarcode(123)
	//item.SetName("giao")
	//item.SetPrice(12.3)
	//err := entityManager.Manager.Save()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//item, err := entityManager.ItemManager.Find(1)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//item, err := entityManager.ItemManager.GetBy("barcode", 123)
	//if err != nil {
	//	log.Fatal(err)
	//}
	item2 := entityManager.ItemManager.New()
	item3 := entityManager.ItemManager.New()
	item2.SetBarcode(128)
	item3.SetBarcode(129)
	item2.SetBelongedItem(item3)
	item3.AddContainedItem(item2)

	//item2 := entityManager.ItemManager.New()
	//item2.SetName("giaoge")
	//err = entityManager.ItemManager.Save(item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if err := entityManager.Manager.Save(); err != nil {
		log.Fatal(err)
	}

	items, _ := item3.GetContainedItem()

	print(items)
}
