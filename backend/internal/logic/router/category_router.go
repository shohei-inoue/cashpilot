package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupCategory はカテゴリルートを登録する
func SetupCategory(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	cc := controller.NewCategoryController(categoryUsecase)

	categories := r.Group("/categories")
	categories.Use(middleware.Auth(jwtSecret))
	{
		categories.GET("", cc.List)
		categories.GET("/:id", cc.Get)
		categories.POST("", cc.Create)
		categories.PUT("/:id", cc.Update)
		categories.DELETE("/:id", cc.Delete)
	}
}
