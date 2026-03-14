package main

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
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
	hc := controller.NewHealthController()
	router.Setup(api, hc)

	r.Run(":8080")
}
