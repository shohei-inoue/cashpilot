package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupUser はユーザールートを登録する
func SetupUser(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	userRepo := repository.NewUserRepository(pool)
	userUsecase := usecase.NewUserUsecase(userRepo)
	uc := controller.NewUserController(userUsecase)

	r.GET("/user", middleware.Auth(jwtSecret), uc.Get)
}
