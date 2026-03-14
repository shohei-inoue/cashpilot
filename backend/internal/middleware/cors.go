package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultAllowOrigin = "http://localhost:3010"

// CORS は CORS ヘッダーを設定するミドルウェア
func CORS() gin.HandlerFunc {
	origin := os.Getenv("CORS_ALLOW_ORIGIN")
	if origin == "" {
		origin = defaultAllowOrigin
	}

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
