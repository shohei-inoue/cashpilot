package middleware

import (
	"net/http"

	"backend/internal/jwt"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

const UserIDKey = "user_id"

// Auth は Cookie の JWT を検証し、user_id をコンテキストに格納するミドルウェア
func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(jwt.CookieName)
		if err != nil || tokenString == "" {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
			c.Abort()
			return
		}

		userID, _, err := jwt.Parse(jwtSecret, tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, response.CodeInvalidToken, response.MsgInvalidToken)
			c.Abort()
			return
		}

		c.Set(UserIDKey, userID)
		c.Next()
	}
}

// GetUserID はコンテキストから user_id を取得する
func GetUserID(c *gin.Context) (int, bool) {
	v, ok := c.Get(UserIDKey)
	if !ok {
		return 0, false
	}
	id, ok := v.(int)
	return id, ok
}

// GetUserIDFromContext はコンテキストから user_id を取得する（ミドルウェア適用済み前提）
func GetUserIDFromContext(c *gin.Context) int {
	id, _ := GetUserID(c)
	return id
}
