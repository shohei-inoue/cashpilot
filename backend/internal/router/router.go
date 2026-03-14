package router

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Setup はルートを登録する。pool が nil の場合は DB 接続チェック・認証ルートをスキップ（テスト用）
func Setup(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	hc := controller.NewHealthController(pool)
	r.GET("/health", hc.Get)

	if pool == nil {
		return
	}

	userRepo := repository.NewUserRepository(pool)
	ac := controller.NewAuthController(jwtSecret, userRepo)
	uc := controller.NewUserController(userRepo)

	// 認証不要
	auth := r.Group("/auth")
	{
		auth.POST("/signup", ac.Signup)
		auth.POST("/login", ac.Login)
		auth.POST("/logout", ac.Logout)
	}

	// 認証必須
	r.GET("/user", middleware.Auth(jwtSecret), uc.Get)
}
