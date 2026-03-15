package router

import (
	"backend/internal/logic/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupHealth はヘルスチェックルートを登録する
func SetupHealth(r *gin.RouterGroup, db *gorm.DB) {
	hc := controller.NewHealthController(db)
	r.GET("/health", hc.Get)
}
