package middleware

import (
	"homework_4_blog/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			// 中断后续中间件的执行
			c.Abort()
			return
		}

		// 2. Bearer Token 格式检查
		// 通常格式: "Bearer <token>"
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token: " + err.Error()})
			c.Abort()
			return
		}

		// 4. 将解析出的声明存储到上下文中，供后续处理函数使用
		c.Set("claims", claims)

		// 5. 继续处理后续的路由处理函数
		c.Next()
	}
}
