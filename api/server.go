package api

import (
	"Cocome/configuration"
	"Cocome/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

var tokenMaker token.Maker

func Start() error {
	config := configuration.Conf
	var err error

	// token 相关
	tokenMaker, err = token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return err
	}

	//获取当前正在使用的Validator，返回的是Interface{}，因此需要类型转换
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("获得Validator错误")
	}

	// 将我们的验证函数注册进来，第一个参数是该验证的自定义名称，使用方法为 binding:"...,currency,..."
	if err := v.RegisterValidation("currency", validCurrency); err != nil {
		return err
	}

	router := gin.Default()
	router.POST("/login", loginUser)
	router.POST("/accounts", createAccount)
	router.GET("/get/:id", getAccount)
	router.GET("/get", listAccount)

	authRoute := router.Group("/auth", authorization)
	authRoute.GET("/", func(c *gin.Context) {
		payload, exist := c.Get(authorizationPayloadKey)
		if !exist {
			c.JSON(http.StatusUnauthorized, errorResponse(errors.New("payload not exists")))
		}
		c.JSON(http.StatusOK, payload)
	})

	return router.Run(config.ServerAddress)
}

//该函数返回一个gin.H，gin.H是一个map，存储着键值对，将要返回给请求者
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
