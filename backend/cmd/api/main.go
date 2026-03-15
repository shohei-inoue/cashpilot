package main

import (
	"log"

	"backend/db"
	"backend/internal/config"
	logicRouter "backend/internal/logic/router"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gormDB, err := db.NewGormDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("failed to get underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()

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
	logicRouter.SetupHealth(api, gormDB)
	logicRouter.SetupAuth(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupUser(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupAccount(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupCategory(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupAnalytics(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupTransaction(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupGoal(api, gormDB, cfg.JWTSecret)
	logicRouter.SetupSimulation(api, gormDB, cfg.JWTSecret)

	r.Run(":8080")
}
