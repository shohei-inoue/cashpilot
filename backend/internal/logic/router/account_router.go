package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAccount は口座ルートを登録する
func SetupAccount(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)
	ac := controller.NewAccountController(accountUsecase)

	accounts := r.Group("/accounts")
	accounts.Use(middleware.Auth(jwtSecret))
	{
		accounts.GET("", ac.List)
		accounts.GET("/:id", ac.Get)
		accounts.POST("", ac.Create)
		accounts.PUT("/:id", ac.Update)
		accounts.DELETE("/:id", ac.Delete)
	}
}
