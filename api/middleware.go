package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("authorization header is not provided")))
		return
	}
	// 前端的authorization 部分应该返回两个字段
	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authorization header format")))
		return
	}
	authType := fields[0]
	if authType != authorizationTypeBearer {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unsupported authorization type %s", authType)))
		return
	}
	token := fields[1]
	payload, err := tokenMaker.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// 用Set设置的键值对在后面的中间件可以用Get来得到
	c.Set(authorizationPayloadKey, payload)
	c.Next()

}
