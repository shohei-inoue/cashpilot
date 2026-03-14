package api

import (
	"context"

	"backend/db"
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

const testJWTSecret = "test-jwt-secret-at-least-32-characters"

// SetupRouterWithDB は DB 接続済みのルーターを返す。接続失敗時は nil, error
func SetupRouterWithDB() (*gin.Engine, error) {
	cfg := config.Load()
	jwtSecret := cfg.JWTSecret
	if jwtSecret == "" {
		jwtSecret = testJWTSecret
	}

	pool, err := db.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	router.Setup(api, pool, jwtSecret)
	return r, nil
}
