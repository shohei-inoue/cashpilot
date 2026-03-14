package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupAnalytics は集計ルートを登録する
func SetupAnalytics(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	analyticsRepo := repository.NewAnalyticsRepository(pool)
	analyticsUsecase := usecase.NewAnalyticsUsecase(analyticsRepo)
	ac := controller.NewAnalyticsController(analyticsUsecase)

	// GET /api/transactions/summary
	transactions := r.Group("/transactions")
	transactions.Use(middleware.Auth(jwtSecret))
	transactions.GET("/summary", ac.Summary)

	// GET /api/analytics/cashflow
	analytics := r.Group("/analytics")
	analytics.Use(middleware.Auth(jwtSecret))
	analytics.GET("/cashflow", ac.Cashflow)
}
