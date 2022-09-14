package main

import (
	"Cocome/entityRepo"
	"Cocome/entityTest"
	"log"
)

func main() {
	//item := entityRepo.ItemRepo.New()
	//item.SetBarcode(123)
	//item.SetName("giao")
	//item.SetPrice(12.3)
	//err := entityRepo.Saver.Save()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//item, err := entityRepo.ItemRepo.Get(1)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//item, err := entityRepo.ItemRepo.GetFromAllInstanceBy("barcode", 123)
	//if err != nil {
	//	log.Fatal(err)
	//}
	item2 := entityTest.ItemRepo.New()
	item3 := entityTest.ItemRepo.New()
	item2.SetBarcode(128)
	item3.SetBarcode(129)
	item2.SetBelongedItem(item3)
	item3.AddContainedItem(item2)
	entityTest.ItemRepo.AddInAllInstance(item2)
	entityTest.ItemRepo.AddInAllInstance(item3)
	//item2 := entityRepo.ItemRepo.New()
	//item2.SetName("giaoge")
	//err = entityRepo.ItemRepo.Save(item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if err := entityRepo.Saver.Save(); err != nil {
		log.Fatal(err)
	}
	item := entityTest.ItemRepo.GetFromAllInstanceBy("barcode", 128)
	print(item)
	if err := entityRepo.Saver.Save(); err != nil {
		log.Fatal(err)
	}

	items := item3.GetContainedItem()
	entityTest.ItemRepo.RemoveFromAllInstance(items[0])

	if err := entityRepo.Saver.Save(); err != nil {
		log.Fatal(err)
	}
	item = entityTest.ItemRepo.GetFromAllInstanceBy("barcode", 128)
	//
	//if err := entityRepo.Saver.Save(); err != nil {
	//	log.Fatal(err)
	//}
	//print(items)
}
