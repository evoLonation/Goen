package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//这个结构体定义将被用来存储http请求的body部分（json）
type transferRequest struct {
	FromAccountId int64 `json:"from_account_id" ` //tag "json"标记该字段对应的json
	ToAccountId   int64 `json:"to_account_id"`
	// gt标记对于realnum也会起作用（0.5之类的）
	Amount   int64  `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof= red green"`
}

func createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if !validAccount(ctx, req.ToAccountId, req.Currency) {
		return
	}
	if !validAccount(ctx, req.FromAccountId, req.Currency) {
		return
	}
	// balabala

	ctx.JSON(http.StatusOK, gin.H{"content": "hello, world!"})
}
func validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	// 执行判断逻辑
	if false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "balabala"})
	}
	return true
}
