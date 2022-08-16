package api

import (
	"Cocome/configuration"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginResponse struct {
	AccessToken string `json:"access_token"`
	UserResponse
}

func loginUser(c *gin.Context) {
	req := &loginRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Username != "evoLonation" || req.Password != "2002116yy" {
		c.JSON(http.StatusNotFound, errorResponse(errors.New("用户不存在或者密码错误")))
		return
	}
	token, err := tokenMaker.CreateToken(req.Username, configuration.Conf.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("用户不存在或者密码错误")))
		return
	}
	res := &loginResponse{
		AccessToken: token,
		UserResponse: UserResponse{
			Username: req.Username,
			Nickname: "赵正阳",
			Email:    "1838940019@qq.com",
		},
	}
	c.JSON(http.StatusOK, res)
	return
}
