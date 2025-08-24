package router

import (
	"homework_4_blog/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API路由分组
	api := router.Group("/api")
	{
		// 用户相关路由
		userGroup := api.Group("/user")
		{
			userGroup.POST("/register", handler.UserRegister)
			userGroup.POST("/login", handler.UserLogin)
		}
	}
	return router
}
