package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupSimulation はシミュレーションルートを登録する
func SetupSimulation(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	analyticsRepo := repository.NewAnalyticsRepository(db)
	goalRepo := repository.NewGoalRepository(db)
	simulationUsecase := usecase.NewSimulationUsecase(analyticsRepo, goalRepo)
	sc := controller.NewSimulationController(simulationUsecase)

	simulation := r.Group("/simulation")
	simulation.Use(middleware.Auth(jwtSecret))
	simulation.POST("/run", sc.Run)
}
