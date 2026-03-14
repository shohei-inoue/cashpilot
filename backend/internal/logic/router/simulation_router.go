package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupSimulation はシミュレーションルートを登録する
func SetupSimulation(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	analyticsRepo := repository.NewAnalyticsRepository(pool)
	goalRepo := repository.NewGoalRepository(pool)
	simulationUsecase := usecase.NewSimulationUsecase(analyticsRepo, goalRepo)
	sc := controller.NewSimulationController(simulationUsecase)

	simulation := r.Group("/simulation")
	simulation.Use(middleware.Auth(jwtSecret))
	simulation.POST("/run", sc.Run)
}
