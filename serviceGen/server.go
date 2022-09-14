package serviceGen

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() error {

	router := gin.Default()

	authRoute := router.Group("/process-sale-service")
	//authRoute.GET("/make-new-sale", makeNewSaleApi())
	authRoute.GET("/enter-item", enterItem)

	return router.Run()
}

//该函数返回一个gin.H，gin.H是一个map，存储着键值对，将要返回给请求者
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//func makeNewSale(ctx *gin.Context) {
//	ret := ProcessSaleServiceInstance.makeNewSale()
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"bool": ret})
//	return
//}

type enterItemRequest struct {
	Barcode  int `json:"barcode"`
	Quantity int `json:"quantity"`
}

func enterItem(ctx *gin.Context) {
	var req enterItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//证明请求对于该结构体并不有效
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ret := ProcessSaleServiceInstance.enterItem(req.Barcode, req.Quantity)
	var errPostCondition *ErrPostCondition
	if errors.Is(ret.Err, ErrPreConditionUnsatisfied) {
		ctx.JSON(http.StatusBadRequest, errorResponse(ret.Err))
		return
	} else if errors.As(ret.Err, &errPostCondition) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.Unwrap(ret.Err)))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bool": ret.Value})
	return
}
