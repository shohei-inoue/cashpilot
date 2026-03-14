package main

import (
	"context"
	"log"

	"backend/db"
	"backend/internal/config"
	"backend/internal/middleware"
	logicRouter "backend/internal/logic/router"

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
	logicRouter.SetupHealth(api, pool)
	logicRouter.SetupAuth(api, pool, cfg.JWTSecret)
	logicRouter.SetupUser(api, pool, cfg.JWTSecret)
	logicRouter.SetupAccount(api, pool, cfg.JWTSecret)
	logicRouter.SetupCategory(api, pool, cfg.JWTSecret)
	logicRouter.SetupTransaction(api, pool, cfg.JWTSecret)

	r.Run(":8080")
}
