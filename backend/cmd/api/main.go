package main

import (
	"context"
	"log"

	"backend/db"
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	r := gin.Default()

	// CORS（credentials 対応のため Allow-Origin は単一オリジン）
	r.Use(middleware.CORS())

	// ルート（疎通確認用）
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CashPilot API",
			"status":  "ok",
		})
	})

	// /api プレフィックス
	api := r.Group("/api")
	router.Setup(api, pool, cfg.JWTSecret)

	r.Run(":8080")
}
