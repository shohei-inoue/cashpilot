package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupAccount は口座ルートを登録する
func SetupAccount(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	accountRepo := repository.NewAccountRepository(pool)
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
