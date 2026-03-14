package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupGoal は目標ルートを登録する
func SetupGoal(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	goalRepo := repository.NewGoalRepository(pool)
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
