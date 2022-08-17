package main

import (
	"Cocome/entityManager"
	"log"
)

func main() {
	item, err := entityManager.ItemManager.Find(1234)
	if err != nil {
		log.Fatal(err)
	}
	item.SetPrice(121233.0)
	item.SetName("giaoge23454744536345")
	item2 := entityManager.ItemManager.Create()
	item2.SetBarcode(12123345)
	item2.SetName("giaoge")
	err = entityManager.ItemManager.Save(item)
	if err != nil {
		log.Fatal(err)
	}
	if err := entityManager.ItemManager.Save(item2); err != nil {
		log.Fatal(err)
	}

	print(item)
}
