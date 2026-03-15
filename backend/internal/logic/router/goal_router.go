package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupGoal は目標ルートを登録する
func SetupGoal(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	goalRepo := repository.NewGoalRepository(db)
	goalUsecase := usecase.NewGoalUsecase(goalRepo)
	gc := controller.NewGoalController(goalUsecase)

	goals := r.Group("/goals")
	goals.Use(middleware.Auth(jwtSecret))
	{
		goals.GET("", gc.List)
		goals.GET("/:id", gc.Get)
		goals.POST("", gc.Create)
		goals.PUT("/:id", gc.Update)
		goals.DELETE("/:id", gc.Delete)
	}
}
