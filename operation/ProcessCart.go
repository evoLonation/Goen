package operation

// CurrentCart 类似于一个微信界面，始终保存购物车对象
//var CurrentCart *entity.Cart
//
//func MakeNewShoppingCart() int {
//	newShoppingCart := entity.Cart{}
//	db.Create(&newShoppingCart)
//	println("create new Shopping Cart, id is " + strconv.Itoa(newShoppingCart.Id))
//	CurrentCart = &newShoppingCart
//	return newShoppingCart.Id
//}
//
//// AddItem 扫码，或者输入条形码的值
//// 这个过程中会更新购物车、商品
//func AddItem(barcode int) (err error) {
//	tx := db.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//			//if .Error != nil {
//			//	err = errors.New("加入购物车失败！")
//			//}
//		}
//		//println(err)
//	}()
//	//锁住购物车
//	if tx.Model(&CurrentCart).Set("gorm:query_option", "FOR UPDATE").Error != nil {
//		return errors.New("给购物车上锁失败！")
//	}
//	var newItem entity.Item
//	//锁住Item
//	if tx.Set("gorm:query_option", "FOR UPDATE").First(&newItem, barcode).Error != nil {
//		return errors.New("没有找到该商品！")
//	}
//	if newItem.StockNumber-1 <= 0 {
//		return errors.New("该商品数量不足！")
//	}
//	newItemInCart := entity.ItemInCart{}
//	newItemInCart.Item = &newItem
//	newItemInCart.Item.StockNumber--
//	newItemInCart.SubAmount = newItem.Price
//	newItemInCart.Quantity = 1
//	CurrentCart.Items = append(CurrentCart.Items, newItemInCart)
//	CurrentCart.Amount += newItemInCart.SubAmount
//
//	if tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&CurrentCart).Error != nil {
//		return errors.New("加入购物车失败！")
//	}
//	return nil
//}
//
//var cartMu sync.Mutex
//
//func AddItemsMul(barcodes []int) (err error) {
//	tx := db.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//	}()
//	//锁住购物车
//	if tx.Model(&CurrentCart).Set("gorm:query_option", "FOR UPDATE").Error != nil {
//		return errors.New("给购物车上锁失败！")
//	}
//	ch := make(chan int, len(barcodes))
//	for i := 0; i < len(barcodes); i++ {
//		go addOneItem(db, barcodes[i], ch)
//	}
//	for i := 0; i < len(barcodes); i++ {
//		<-ch
//	}
//	if tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&CurrentCart).Error != nil {
//		return errors.New("加入购物车失败！")
//	}
//	return nil
//}
//func addOneItem(db *gorm.DB, barcode int, ch chan int) (err error) {
//	println("我开始了")
//	tx := db.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//		ch <- 0
//	}()
//	var newItem entity.Item
//	//锁住Item
//	if tx.Set("gorm:query_option", "FOR UPDATE").First(&newItem, barcode).Error != nil {
//		return errors.New("没有找到该商品！")
//	}
//	if newItem.StockNumber-1 <= 0 {
//		return errors.New("该商品数量不足！")
//	}
//	arr := [1000000]int{}
//	for i := 0; i < 1000000; i++ {
//		for j := 0; j < 100; j++ {
//			arr[i] = i + j*114514
//		}
//	}
//	newItemInCart := entity.ItemInCart{}
//	newItemInCart.Item = &newItem
//	newItemInCart.Item.StockNumber--
//	newItemInCart.SubAmount = newItem.Price
//	newItemInCart.Quantity = 1
//	cartMu.Lock()
//	CurrentCart.Items = append(CurrentCart.Items, newItemInCart)
//	CurrentCart.Amount += newItemInCart.SubAmount
//	cartMu.Unlock()
//	println("我结束了")
//	return nil
//}
//
//func AddItems(barcodes []int) (err error) {
//	tx := db.Begin()
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//	}()
//	//锁住购物车
//	if tx.Model(&CurrentCart).Set("gorm:query_option", "FOR UPDATE").Error != nil {
//		return errors.New("给购物车上锁失败！")
//	}
//	for i := 0; i < len(barcodes); i++ {
//		println("我开始了")
//		barcode := barcodes[i]
//		var newItem entity.Item
//		//锁住Item
//		if tx.Set("gorm:query_option", "FOR UPDATE").First(&newItem, barcode).Error != nil {
//			return errors.New("没有找到该商品！")
//		}
//		if newItem.StockNumber-1 <= 0 {
//			return errors.New("该商品数量不足！")
//		}
//		arr := [1000000]int{}
//		for i := 0; i < 1000000; i++ {
//			for j := 0; j < 100; j++ {
//				arr[i] = i + j*114514
//			}
//		}
//		newItemInCart := entity.ItemInCart{}
//		newItemInCart.Item = &newItem
//		newItemInCart.Item.StockNumber--
//		newItemInCart.SubAmount = newItem.Price
//		newItemInCart.Quantity = 1
//		CurrentCart.Items = append(CurrentCart.Items, newItemInCart)
//		CurrentCart.Amount += newItemInCart.SubAmount
//		println("我结束了")
//	}
//	if tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&CurrentCart).Error != nil {
//		return errors.New("加入购物车失败！")
//	}
//	return nil
//}
//
//func AddNumber(i int) error {
//	itemInShopping := CurrentCart.Items[i]
//	if itemInShopping.Item.StockNumber == 0 {
//		return errors.New("商品数量不足！")
//	}
//	itemInShopping.Item.StockNumber--
//	itemInShopping.Quantity++
//	itemInShopping.SubAmount += itemInShopping.Item.Price
//	if db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&itemInShopping).Error != nil {
//		return errors.New("增加数量失败！")
//	}
//	return nil
//}

//func EndSale() float64 {
//	var salesLineItems []entity.SalesLineItem
//	db.Model(&currentSale).Association("SalesLineItems").Find(&salesLineItems)
//	for i := 0; i < len(salesLineItems); i++ {
//		currentSale.Amount += salesLineItems[i].SubAmount
//	}
//	currentSale.IsReadyToPay = true
//	return currentSale.Amount
//}
//
//func MakeCashPayment(amount float64) bool {
//	if currentSale.Amount > amount {
//		return false
//	}
//	payment := entity.Payment{
//		AmountTendered: amount,
//		Balance:        amount - currentSale.Amount,
//	}
//	currentSale.Payment = payment
//	currentSale.IsComplete = true
//	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&currentSale)
//	return true
//}
