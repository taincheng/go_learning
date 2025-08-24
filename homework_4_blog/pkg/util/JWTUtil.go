package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT配置
var (
	// 密匙，生产环境用环境变量
	jwtSecret = []byte("XXXXXXXXXXXX")
	// 过期时间
	tokenExpireDuration = time.Hour * 24
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	// jwt声明结构体
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
			Issuer:    "gin-blog",                                              // 签发者
			Subject:   "user.token",                                            // 主题
		},
	}
	// 创建 Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名并生成完整的编码Token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 2. 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// 3. 返回密钥用于验证签名
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 4. 验证 Token 是否有效并获取声明
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
