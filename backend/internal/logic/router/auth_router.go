package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupAuth は認証ルートを登録する
func SetupAuth(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	userRepo := repository.NewUserRepository(pool)
	authUsecase := usecase.NewAuthUsecase(jwtSecret, userRepo)
	ac := controller.NewAuthController(authUsecase)

	auth := r.Group("/auth")
	{
		auth.POST("/signup", ac.Signup)
		auth.POST("/login", ac.Login)
		auth.POST("/logout", ac.Logout)
	}
}
