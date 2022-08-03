package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//这个结构体定义将被用来存储http请求的body部分（json）
type createAccountRequest struct {
	Owner    string `json:"owner" ` //tag "json"标记该字段对应的json
	Currency string `json:"currency" binding:"required,oneof= USD RMB"`
	// binding required 标记该字段是必须的 oneof 标记该字段只能是规定的某几种
}

func createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//证明请求对于该结构体并不有效
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// 在之后的处理部分，也可以使用上面的方式来返回错误
	ctx.JSON(http.StatusOK, gin.H{"content": "hello, world!"})
	return
}

type getAccountRequest struct {
	Id int64 `uri:"id" binding:"min=1"`
}

func getAccount(ctx *gin.Context) {
	var req getAccountRequest
	// 绑定所有 url中的参数
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 在之后的处理部分，也可以使用上面的方式来返回错误
	ctx.JSON(http.StatusOK, gin.H{"content": "hello, world!, 你获得了 id为" + strconv.Itoa(int(req.Id))})
}

type listAccountRequest struct {
	PageId   int32 `form:"page_id" binding:"min=1"`   //tag "json"标记该字段对应的json
	PageSize int32 `form:"page_size" binding:"min=1"` //tag "json"标记该字段对应的json
}

func listAccount(ctx *gin.Context) {
	var req listAccountRequest
	// 绑定 sql params
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// 在之后的处理部分，也可以使用上面的方式来返回错误
	ctx.JSON(http.StatusOK, gin.H{"content": "hello, world!, 你获得了 id为" + strconv.Itoa(int(req.PageId))})
}
