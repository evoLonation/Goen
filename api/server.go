package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start(address string) {
	router := gin.Default()

	//获取当前正在使用的Validator，返回的是Interface{}，因此需要类型转换
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		print("获得Validator错误")
	}

	// 将我们的验证函数注册进来，第一个参数是该验证的自定义名称，使用方法为 binding:"...,currency,..."
	v.RegisterValidation("currency", validCurrency)

	router.POST("/accounts", createAccount)
	router.GET("/get/:id", getAccount)
	router.GET("/get", listAccount)

	router.Run(address)
}

//该函数返回一个gin.H，gin.H是一个map，存储着键值对，将要返回给请求者
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
