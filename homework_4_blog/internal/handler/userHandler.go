package handler

import (
	"encoding/json"
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/service"
	"homework_4_blog/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
//使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var user model.User

	var userData string
	if userJson, jsonErr := json.Marshal(user); jsonErr == nil {
		userData = string(userJson)
	}

	// 接受注册数据
	if err := c.ShouldBindJSON(&user); err != nil {

		util.Logger.Error(
			"绑定注册数据失败",
			zap.String("errMessage", err.Error()),
			zap.String("userData", userData),
		)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码，存储用户数据
	err1 := service.CreateUser(&user)
	if err1 != nil {
		util.Logger.Error("注册用户失败", zap.String("errMessage", err1.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.SelectUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	// 将 token 放在响应头中
	c.Header("Authorization", "Bearer "+*token)

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   *token,
	})
}
