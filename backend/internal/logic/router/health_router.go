package router

import (
	"backend/internal/logic/controller"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupHealth はヘルスチェックルートを登録する
func SetupHealth(r *gin.RouterGroup, pool *pgxpool.Pool) {
	hc := controller.NewHealthController(pool)
	r.GET("/health", hc.Get)
}
