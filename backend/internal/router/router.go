package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

// Setup はルートを登録する
func Setup(r *gin.RouterGroup, hc *controller.HealthController) {
	// ヘルスチェック（認証不要）
	r.GET("/health", hc.Get)

	// 今後追加:
	// AuthRouter(r.Group("/auth"), ac)
	// AccountRouter(r.Group("/accounts"), acc)
	// ...
}
