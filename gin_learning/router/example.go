package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 开启服务器的简单例子
func startExample() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

// Restful风格接口规范
func restfulExample() {
	router := gin.Default()
	router.GET("/ping1")    // 查询
	router.POST("/ping2")   // 新增
	router.DELETE("/ping3") // 删除
	router.PUT("/ping4")    // 更新（客户端提供完整数据）
	router.PATCH("/ping5")  // 更新（客户端提供需要修改的数据）

	router.Run()
}

// 路由分组
func routerGroup() {
	router := gin.Default()
	// v1 分组
	v1 := router.Group("v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// v2 分组
	v2 := router.Group("v2")
	v2.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}

// 重定向
func redirect() {
	router := gin.Default()

	router.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"foo": "foo",
		})
	})

	// 重定向到外部
	router.GET("/redirect1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// 重定向到内部
	router.GET("/redirect2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	// 修改request路由，实现重定向访问，实际上是只访问了test2
	router.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})

	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	router.Run()
}

// 访问静态文件
func getStaticFile() {
	router := gin.Default()
	// assets必须在项目根目录下，只能访问指定的文件
	// http://localhost:8080/assets/test1.txt
	//router.Static("/assets", "./assets")

	// 可以访问目录，列出文件
	// http://localhost:8080/assets/
	//router.StaticFS("/assets", http.Dir("./assets"))

	// 访问单个文件
	// http://localhost:8080/test1.txt
	router.StaticFile("/test1", "./assets/test1.txt") // 单独的文件
	router.Run()
}

// 向浏览器输出
func outPut() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// 有内置的多种输出格式
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		//c.XML(200, gin.H{
		//	"message": "hello world",
		//})
		//c.String(200, "hello world")
	})
	router.Run()
}
func Run() {
	outPut()
}
