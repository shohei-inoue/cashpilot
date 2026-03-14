package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HealthController はヘルスチェック用のコントローラー
type HealthController struct {
	pool *pgxpool.Pool
}

// NewHealthController は HealthController を生成する
func NewHealthController(pool *pgxpool.Pool) *HealthController {
	return &HealthController{pool: pool}
}

// Get は GET /api/health のハンドラー。DB 接続確認を含む（pool が nil の場合はスキップ）
func (h *HealthController) Get(c *gin.Context) {
	if h.pool != nil {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := h.pool.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  "database unreachable",
			})
			return
		}
	}


	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
