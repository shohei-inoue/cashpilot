package controller

import (
	"context"
	"net/http"
	"time"

	"backend/internal/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthController はヘルスチェック用のコントローラー
type HealthController struct {
	db *gorm.DB
}

// NewHealthController は HealthController を生成する
func NewHealthController(db *gorm.DB) *HealthController {
	return &HealthController{db: db}
}

// Get は GET /api/health のハンドラー。DB 接続確認を含む（db が nil の場合はスキップ）
func (h *HealthController) Get(c *gin.Context) {
	if h.db != nil {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		sqlDB, err := h.db.DB()
		if err != nil {
			response.Error(c, http.StatusServiceUnavailable, response.CodeDatabaseError, response.MsgDatabaseUnreachable)
			return
		}
		if err := sqlDB.PingContext(ctx); err != nil {
			response.Error(c, http.StatusServiceUnavailable, response.CodeDatabaseError, response.MsgDatabaseUnreachable)
			return
		}
	}

	response.Success(c, gin.H{response.MsgStatus: response.MsgOK})
}
