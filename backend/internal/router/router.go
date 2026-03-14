package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Setup はルートを登録する。pool が nil の場合は DB 接続チェックをスキップ（テスト用）
func Setup(r *gin.RouterGroup, pool *pgxpool.Pool) {
	hc := controller.NewHealthController(pool)
	// ヘルスチェック（認証不要）
	r.GET("/health", hc.Get)

	// 今後追加:
	// AuthRouter(r.Group("/auth"), ac)
	// AccountRouter(r.Group("/accounts"), acc)
	// ...
}
