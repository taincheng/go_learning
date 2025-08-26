package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

type FormA struct {
	UserName string `json:"u" xml:"u" query:"u" form:"u" binding:"required"`
	Password string `json:"p" xml:"p" query:"p" form:"p"`
}

type FormB struct {
	Job string `xml:"job" json:"job" `
}

// api 参数绑定
func paramBind() {
	router := gin.Default()
	router.POST("/bind", func(c *gin.Context) {
		var formA FormA
		var formB FormB

		// ShouldBind 自动绑定符合配置格式的参数
		//if err := c.ShouldBind(&formA); err == nil {
		//	fmt.Println(formA)
		//	c.JSON(200, gin.H{
		//		"username": formA.UserName,
		//		"password": formA.Password,
		//	})
		//} else {
		//	fmt.Println(err)
		//}

		// c.ShouldBind 使用了 c.Request.Body，不可重用。应该使用 ShouldBindWith
		// Key: 'FormA.UserName' Error:Field validation for 'UserName' failed on the 'required' tag EOF
		//if errA := c.ShouldBind(&formA); errA == nil {
		//	c.String(http.StatusOK, `the body should be formA`)
		//	// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		//} else if errB := c.ShouldBind(&formB); errB == nil {
		//	c.String(http.StatusOK, `the body should be formB`)
		//} else {
		//	fmt.Println(errA, errB)
		//}

		// 读取 c.Request.Body 并将结果存入上下文。 ShouldBindBodyWith 多次绑定
		if errA := c.ShouldBindBodyWith(&formA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 这时, 复用存储在上下文中的 body。可以使用其他的格式
		} else if errB := c.ShouldBindBodyWith(&formB, binding.XML); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
			if errA != nil {
				fmt.Printf("errA: %v", errA)
			}
			if errB != nil {
				fmt.Printf("errB: %v", errB)
			}
		}

	})
	router.Run()
}

func handler1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("handler1 before")
		c.Next()
		fmt.Println("handler1 after")
	}
}

func handler2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("handler2 before")
		c.Next()
		fmt.Println("handler2 after")
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthRequired")
		c.Next()
	}
}

// 自定义中间件
func handlerTest() {
	router := gin.Default()
	router.GET("/handler", handler1(), handler2(), func(c *gin.Context) {
		fmt.Println("self")
		c.String(http.StatusOK, "self handler test")
	})
	router.Run()
}

// 使用中间件
func useHandler() {
	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	router.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	router.GET("/benchmark", handler1(), handler2())

	// 认证路由组
	// authorized := router.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := router.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", handler1())
		authorized.POST("/submit", handler1())
		authorized.POST("/read", handler1())

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", handler1())
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// 用户认证
func authorizedTest() {
	router := gin.Default()

	// gin.BasicAuth() 是一个中间件，为整个路由组添加HTTP基本认证
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// MustGet 从上下文中获取认证用户信息
		// .(string) 类型断言为 string，失败报错
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	router.Run()
}

type LoginInfo struct {
	Username string `json:"username" form:"username" binding:"required"` // 通过 binding 声明校验规则
	Password string `json:"password" form:"password" binding:"number,eq=1111"`
	Email    string `json:"email" form:"email" binding:"email"`
}

// 验证器，gin 使用 github.com/go-playground/validator/v10
func validatorTest() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		login := LoginInfo{}
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, login)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// MyStruct ..
type MyStruct struct {
	Str string `json:"str" form:"str" validate:"is-awesome"`
}

// validateMyVal implements validator.Func 自定义验证器规则
func validateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}

// 自定义验证器的使用
func myValidatorTest() {
	router := gin.Default()
	validate := validator.New()
	err := validate.RegisterValidation("is-awesome", validateMyVal)
	if err != nil {
		return
	}
	myStruct := MyStruct{}
	router.POST("/", func(c *gin.Context) {
		c.ShouldBind(&myStruct)
		if err2 := validate.Struct(myStruct); err2 != nil {
			c.JSON(http.StatusOK, gin.H{"error": err2.Error(), "info": "字符串不是 is-awesome"})
		} else {
			c.JSON(http.StatusOK, gin.H{"info": "验证通过"})
		}
	})
	router.Run()
}

func Run() {
	validatorTest()
}
