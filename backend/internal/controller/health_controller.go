package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController はヘルスチェック用のコントローラー
type HealthController struct{}

// NewHealthController は HealthController を生成する
func NewHealthController() *HealthController {
	return &HealthController{}
}

// Get は GET /api/health のハンドラー
func (h *HealthController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
