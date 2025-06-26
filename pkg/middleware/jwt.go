package middleware

import (
	"errors"
	"os"
	"strings"
	"time"

	"erp-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "your-256-bit-secret" // 默认密钥，建议在生产环境中通过环境变量设置
	}
	return secret
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, userType string) (string, error) {
	claims := Claims{
		UserID:   userID,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时后过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.UnauthorizedResponse(c, "未提供认证信息")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.UnauthorizedResponse(c, "认证格式错误")
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("无效的签名方法")
			}
			return jwtSecret, nil
		})

		if err != nil {
			response.UnauthorizedResponse(c, "无效的令牌")
			c.Abort()
			return
		}

		if !token.Valid {
			response.UnauthorizedResponse(c, "令牌已过期")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.UserType)
		c.Next()
	}
}

// RequireUserType 检查用户类型的中间件
func RequireUserType(allowedTypes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("user_type")
		if !exists {
			response.UnauthorizedResponse(c, "未找到用户类型信息")
			c.Abort()
			return
		}

		userTypeStr := userType.(string)
		allowed := false
		for _, t := range allowedTypes {
			if userTypeStr == t {
				allowed = true
				break
			}
		}

		if !allowed {
			response.ForbiddenResponse(c, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}
