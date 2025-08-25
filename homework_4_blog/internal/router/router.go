package router

import (
	"homework_4_blog/internal/handler"
	"homework_4_blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API路由分组
	api := router.Group("/api")
	api.Use(middleware.LoggerMiddleware())
	{
		// 用户相关路由
		userGroup := api.Group("/user")
		{
			userGroup.POST("/register", handler.UserRegister)
			userGroup.POST("/login", handler.UserLogin)
		}
		postGroup := api.Group("/post")
		postGroup.Use(middleware.AuthMiddleware())
		{
			postGroup.POST("/createPost", handler.CreatePost)
			postGroup.GET("/getPostList", handler.GetPostList)
			postGroup.GET("/getPostInfo", handler.GetPostInfo)
			postGroup.POST("/updatePost", handler.UpdatePost)
			postGroup.POST("/deletePost", handler.DeletePost)
		}
	}
	return router
}
