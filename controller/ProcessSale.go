package controller

//func Start() {
//	r := gin.Default()
//	r.POST("/add", AddInCart)
//	r.POST("/addsm", AddsInCartMul)
//	r.POST("/adds", AddsInCart)
//	r.Run()
//}

//func AddInCart(c *gin.Context) {
//	barcode, _ := strconv.Atoi(c.sql("id"))
//	if operation.AddItem(barcode) != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "error",
//			"message": c.DefaultPostForm("nick", "anonymous"),
//		})
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"status":  "posted",
//		"message": "成功",
//	})
//}
//
//type Request struct {
//	Ids []int
//}
//
//func AddsInCartMul(c *gin.Context) {
//	request := Request{}
//	c.BindJSON(&request)
//	if operation.AddItemsMul(request.Ids) != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "error",
//			"message": c.DefaultPostForm("nick", "anonymous"),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "posted",
//			"message": "成功",
//		})
//	}
//}
//func AddsInCart(c *gin.Context) {
//	request := Request{}
//	c.BindJSON(&request)
//	if operation.AddItems(request.Ids) != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "error",
//			"message": c.DefaultPostForm("nick", "anonymous"),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "posted",
//			"message": "成功",
//		})
//	}
//}
